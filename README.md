# caas-api-gateway
📦 Service Bindings (for local dev)
Service	Port	Notes
```
caas-eks-api-go	:8081	Run with make run from EKS repo
caas-aks-api-go	:8082	Run with make run from AKS repo
caas-api-gateway	:8080	Unified access point
```
Update their main.go as needed to listen on :8081 and :8082.

🌐 Access Points
API Path	Target
```
/aws/clusters	Proxies to EKS API
/azure/clusters	Proxies to AKS API
/swagger/aws/	Swagger for EKS
/swagger/azure/	Swagger for AKS
```
/	You can add a root UI

🔧 Optional: Merged Swagger UI
If you want to merge both Swagger definitions:

Export EKS and AKS Swagger as JSON:

```
curl http://localhost:8081/swagger/doc.json > eks.json
curl http://localhost:8082/swagger/doc.json > aks.json
Use a tool like Swagger Combine or SwaggerHub to merge:
```
```
apis:
  - url: ./eks.json
  - url: ./aks.json
```
Serve merged JSON via /swagger/combined.json and use Swagger UI to load it.
