# AWSSDKPluginGo

A simple Golang-based Container-as-a-Service (CaaS) REST API that manages AWS EKS clusters using the AWS SDK for Go (v2).  
It allows you to create, delete, list clusters and deploy Kubernetes manifests via `kubectl`.

---

## Directory Structure
```
caas-eks/
├── main.go # Main API server and EKS logic
├── go.mod # Go module dependencies
├── deployment.yaml # Kubernetes manifest to deploy to EKS
---
```
## Features

- Create EKS Clusters  
- Delete EKS Clusters  
- List existing clusters  
- Describe specific cluster  
- Deploy Kubernetes apps using `kubectl apply`

---

## Prerequisites

- [Go 1.19+](https://golang.org/doc/install)
- AWS CLI configured (`aws configure`)
- `kubectl` installed and in your system `PATH`
- IAM permissions for EKS + EC2 + IAM + VPC

---

## Setup

```bash
git clone https://github.com/mpwusr/AWSSDKPluginGo.git
cd caas-eks
go mod tidy
go run main.go
```
API Usage
Create a Cluster
```
curl -X POST http://localhost:8080/clusters \
  -H "Content-Type: application/json" \
  -d '{
    "name": "caas-demo",
    "role_arn": "arn:aws:iam::123456789012:role/EKSClusterRole",
    "subnet_ids": ["subnet-abc", "subnet-def"],
    "security_groups": ["sg-01234"],
    "version": "1.27"
  }'
```
List All Clusters
```
curl http://localhost:8080/clusters
```
Deploy an App to a Cluster
Ensure deployment.yaml is valid and present.
```
curl -X POST http://localhost:8080/clusters/caas-demo/deploy
```
License
MIT License. See LICENSE for more info.





