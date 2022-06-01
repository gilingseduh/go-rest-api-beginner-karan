package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api-beginner-karan/pkg/mocks"
	"go-rest-api-beginner-karan/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	for index, book := range mocks.Books {
		if book.Id == id {
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(fmt.Sprintf("Book with ID %d not found", id))
}
