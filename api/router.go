package api

import (
	"caas-eks-api-go/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/clusters", handlers.CreateCluster).Methods("POST")
	r.HandleFunc("/clusters", handlers.ListClusters).Methods("GET")
	r.HandleFunc("/clusters/{name}", handlers.GetCluster).Methods("GET")
	r.HandleFunc("/clusters/{name}", handlers.DeleteCluster).Methods("DELETE")
	r.HandleFunc("/clusters/{name}/deploy", handlers.DeployApp).Methods("POST")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r
}
