package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-api-beginner-karan/pkg/mocks"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
