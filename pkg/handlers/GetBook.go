package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-api-beginner-karan/pkg/mocks"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books {
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
