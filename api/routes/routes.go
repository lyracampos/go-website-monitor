package routes

import (
	"log"
	"net/http"

	"website-monitor/api/controllers"

	"github.com/gorilla/mux"
)

func Load() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/websites", controllers.List).Methods(http.MethodGet)
	api.HandleFunc("/websites/{websiteId}", controllers.Get).Methods(http.MethodGet)
	api.HandleFunc("/websites", controllers.Add).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
