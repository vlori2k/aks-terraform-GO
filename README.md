
# AKS Terraform to Go 🚀

This project demonstrates how to provision and manage Azure infrastructure (Resource Group and AKS Cluster) using **Go** and the Azure SDK.  
It mirrors an existing Terraform setup and shows how the same tasks can be performed programmatically with Go.

---

## 📌 Contents
- `main.go` – Go code to create a Resource Group and AKS Cluster in Azure.
- `go.mod` / `go.sum` – Go module and dependency files.
- `docs/Dokumentasjon Norsk - komme igang med GO og Azure.docx` – Step-by-step documentation (in Norwegian).

---

## ⚙️ Prerequisites
- [Go](https://go.dev/dl/) installed  
- [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) installed  
- An active **Azure Subscription**  
- Log in to Azure:
  ```bash
  az login

🚀 Running the project

Clone the repo:

git clone https://github.com/<your-username>/aks-terraform-go.git
cd aks-terraform-go


Fetch dependencies:

go mod tidy


Run the code:

go run main.go


This will provision:

A Resource Group (aks-demo-rg2)

An AKS Cluster (aks-demo-cluster2)



📚 Documentation

Full step-by-step documentation (in Norwegian) can be found here:
➡️ Dokumentasjon Norsk - komme igang med GO og Azure.docx

💡 Why this project?

Demonstrates how the same infrastructure can be provisioned with Terraform or Go.

Provides hands-on experience with the Azure SDK for Go.

Useful as a demo project to showcase at interviews.