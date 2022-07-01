package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type API struct{}

var books = []string{"Book 1", "Book 2", "Book 3"}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	limitBooks := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitBooks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if limit < 0 || limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[:limit])
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request) {}
