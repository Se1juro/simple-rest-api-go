package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Se1juro/simple-rest-api-go/api"
	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\":\"Hello World\"}")
}

func main() {
	router := mux.NewRouter()
	// API Object
	a := &api.API{}
	a.RegisterRoutes(router)
	// EndPoints
	router.HandleFunc("/api/index", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":3001",
		Handler: router,
	}
	log.Println("Listening on port 3001")
	srv.ListenAndServe()
}
