package handlers

import (
	"encoding/json"
	"go-rest-api-beginner-karan/pkg/mocks"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mocks.Books)
}
