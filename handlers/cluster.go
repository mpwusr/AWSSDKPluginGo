package handlers

import (
	"caas-eks-api-go/models"
	"caas-eks-api-go/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary Create EKS Cluster
// @Tags EKS
// @Accept json
// @Produce json
// @Param cluster body models.CreateClusterRequest true "Cluster config"
// @Success 200 {object} interface{}
// @Failure 500 {object} map[string]string
// @Router /clusters [post]
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	var req models.CreateClusterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", 400)
		return
	}
	cluster, err := service.CreateCluster(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("create error: %v", err), 500)
		return
	}
	json.NewEncoder(w).Encode(cluster)
}

func ListClusters(w http.ResponseWriter, _ *http.Request) {
	clusters, err := service.ListClusters()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(clusters)
}

func GetCluster(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	cluster, err := service.GetCluster(name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(cluster)
}

func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := service.DeleteCluster(name); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("Cluster deletion started"))
}

func DeployApp(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := service.DeployApp(name); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("App deployed"))
}
