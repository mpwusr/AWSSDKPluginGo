# AWSSDKPluginGo

A simple Golang-based Container-as-a-Service (CaaS) REST API that manages AWS EKS clusters using the AWS SDK for Go (v2).  
It allows you to create, delete, list clusters and deploy Kubernetes manifests via `kubectl`.

---

## 📁 Directory Structure

caas-eks/
├── main.go # Main API server and EKS logic
├── go.mod # Go module dependencies
├── deployment.yaml # Kubernetes manifest to deploy to EKS
---

## Features

- Create EKS Clusters  
- Delete EKS Clusters  
- List existing clusters  
- Describe specific cluster  
- Deploy Kubernetes apps using `kubectl apply`

---

## 🔧 Prerequisites

- [Go 1.19+](https://golang.org/doc/install)
- AWS CLI configured (`aws configure`)
- `kubectl` installed and in your system `PATH`
- IAM permissions for EKS + EC2 + IAM + VPC

---

## 🛠️ Setup

```bash
git clone https://github.com/yourusername/AWSSDKPluginGo.git
cd caas-eks
go mod tidy
go run main.go

