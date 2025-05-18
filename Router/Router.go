package Router

import (
	SalseTracker "SalesTracker/SalseTracker"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(pRouter *mux.Router) {
	log.Println("RegisterRoutes (-)")
	pRouter.HandleFunc("/SyncData", SalseTracker.SyncDataApi).Methods(http.MethodGet)
	pRouter.HandleFunc("/getRevenue/{id}", SalseTracker.GetRevenueApi).Methods(http.MethodGet)
	log.Println("RegisterRoutes (-)")
}