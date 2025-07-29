# AWSSDKPluginGo

Directory structure

caas-eks/
├── main.go
├── go.mod
├── deployment.yaml       <-- your k8s manifest

# Create a cluster
curl -X POST http://localhost:8080/clusters -H "Content-Type: application/json" -d '{
  "name": "caas-demo",
  "role_arn": "arn:aws:iam::123456789012:role/EKSClusterRole",
  "subnet_ids": ["subnet-abc", "subnet-def"],
  "security_groups": ["sg-01234"],
  "version": "1.27"
}'

# List clusters
curl http://localhost:8080/clusters

# Deploy app
curl -X POST http://localhost:8080/clusters/caas-demo/deploy

