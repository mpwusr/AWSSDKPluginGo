package main

import (
	"log"
	"net/http"

	"caas-eks-api-go/api"
)

// @title CaaS EKS API
// @version 1.0
// @description API to manage AWS EKS clusters
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	r := api.SetupRouter()
	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
