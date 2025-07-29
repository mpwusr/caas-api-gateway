package router

import (
	"caas-api-gateway/proxy"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()

	// Reverse proxies to individual services
	r.PathPrefix("/aws/").Handler(proxy.NewReverseProxy("http://localhost:8081")) // EKS API
	r.PathPrefix("/azure/").Handler(proxy.NewReverseProxy("http://localhost:8082")) // AKS API

	// Swagger UI reverse proxies
	r.PathPrefix("/swagger/aws/").Handler(proxy.NewReverseProxy("http://localhost:8081/swagger/"))
	r.PathPrefix("/swagger/azure/").Handler(proxy.NewReverseProxy("http://localhost:8082/swagger/"))

	return r
}
