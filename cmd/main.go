package main

import (
	"log"
	"net/http"

	"caas-api-gateway/router"
)

func main() {
	r := router.Setup()
	log.Println("🚀 CaaS Gateway running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
