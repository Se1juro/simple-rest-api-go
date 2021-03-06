package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/books", a.getBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/books/{id}", a.getBook).Methods(http.MethodGet)
}
