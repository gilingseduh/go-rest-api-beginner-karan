package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api-beginner-karan/pkg/models"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(books)
}
