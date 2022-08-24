package main

import (
	"awesomeProject1/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting at port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router.Limit(r)))
}
