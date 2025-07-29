
---

## 🚀 Features

- ✅ Create EKS Clusters  
- ❌ Delete EKS Clusters  
- 🔁 List existing clusters  
- 🔍 Describe specific cluster  
- 🚀 Deploy Kubernetes apps using `kubectl apply`

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

