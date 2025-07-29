package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/gorilla/mux"
)

// Global AWS EKS client
var eksClient *eks.Client

func main() {
	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}
	eksClient = eks.NewFromConfig(cfg)

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/clusters", createCluster).Methods("POST")
	r.HandleFunc("/clusters", listClusters).Methods("GET")
	r.HandleFunc("/clusters/{name}", getCluster).Methods("GET")
	r.HandleFunc("/clusters/{name}", deleteCluster).Methods("DELETE")
	r.HandleFunc("/clusters/{name}/deploy", deployApp).Methods("POST")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Request payload for cluster creation
type CreateClusterRequest struct {
	Name       string   `json:"name"`
	RoleArn    string   `json:"role_arn"`
	SubnetIds  []string `json:"subnet_ids"`
	SecurityGs []string `json:"security_groups"`
	Version    string   `json:"version"`
}

// POST /clusters
func createCluster(w http.ResponseWriter, r *http.Request) {
	var req CreateClusterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	input := &eks.CreateClusterInput{
		Name:    aws.String(req.Name),
		RoleArn: aws.String(req.RoleArn),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			SubnetIds:        req.SubnetIds,
			SecurityGroupIds: req.SecurityGs,
		},
		Version: aws.String(req.Version),
	}

	resp, err := eksClient.CreateCluster(context.TODO(), input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create cluster: %v", err), 500)
		return
	}

	json.NewEncoder(w).Encode(resp.Cluster)
}

// GET /clusters
func listClusters(w http.ResponseWriter, r *http.Request) {
	resp, err := eksClient.ListClusters(context.TODO(), &eks.ListClustersInput{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list clusters: %v", err), 500)
		return
	}

	json.NewEncoder(w).Encode(resp.Clusters)
}

// GET /clusters/{name}
func getCluster(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	resp, err := eksClient.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{
		Name: aws.String(name),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to describe cluster: %v", err), 500)
		return
	}

	json.NewEncoder(w).Encode(resp.Cluster)
}

// DELETE /clusters/{name}
func deleteCluster(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	_, err := eksClient.DeleteCluster(context.TODO(), &eks.DeleteClusterInput{
		Name: aws.String(name),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete cluster: %v", err), 500)
		return
	}
	w.Write([]byte("Cluster deletion started"))
}

// POST /clusters/{name}/deploy
func deployApp(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	kubeconfigPath := fmt.Sprintf("/tmp/%s-kubeconfig.yaml", name)

	// Update kubeconfig (requires AWS CLI installed)
	cmd := exec.Command("aws", "eks", "update-kubeconfig", "--name", name, "--kubeconfig", kubeconfigPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update kubeconfig: %v\n%s", err, string(out)), 500)
		return
	}

	// Apply manifest (hardcoded YAML path for now)
	kubectl := exec.Command("kubectl", "apply", "-f", "deployment.yaml", "--kubeconfig", kubeconfigPath)
	if out, err := kubectl.CombinedOutput(); err != nil {
		http.Error(w, fmt.Sprintf("kubectl apply failed: %v\n%s", err, string(out)), 500)
		return
	}

	w.Write([]byte("Application deployed to EKS"))
}
