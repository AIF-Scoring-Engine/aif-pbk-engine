package router

import (
	"awesomeProject1/controller"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"log"
	"net"
	"net/http"
	"sync"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 4000)
		visitors[ip] = limiter
	}

	return limiter
}

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the IP address for the current user.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Call the getVisitor function to retreive the rate limiter for
		// the current user.
		limiter := getVisitor(ip)
		fmt.Println(ip)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429)+ip, http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/post/company", controller.PostCompany).Methods("POST")
	router.HandleFunc("/api/post/dev/company", controller.PostCompanyDev).Methods("POST")

	return router
}
