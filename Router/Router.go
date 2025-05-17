package Router

import (
	// integration "SalesAnalytics/Integration"
	// "SalesAnalytics/handlers"

	Api "SalesTracker/Task"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(pRouter *mux.Router) {
		log.Println("RegisterRoutes (-)")

	// pRouter.HandleFunc("/api/refresh", Api.DemoApi).Methods(http.MethodGet)
	pRouter.HandleFunc("/api/refresh/{id}", Api.DemoApi).Methods(http.MethodGet)
	// r.HandleFunc("/api/revenue/{id}", handlers.RevenueHandler)
	// r.HandleFunc("/api/nProducts/{id}", integration.TopNProducts)
	// r.HandleFunc("/api/customers/{id}", integration.CustomerAnalysis)
	// r.HandleFunc("/api/calculations/{id}", integration.SalesCalculations)
		log.Println("RegisterRoutes (-)")

}