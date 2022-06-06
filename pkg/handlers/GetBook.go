package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api-beginner-karan/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(fmt.Sprintf("Book with ID %d not found", id))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&book)
	}
}
