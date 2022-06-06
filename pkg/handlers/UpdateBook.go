package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api-beginner-karan/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var existingBook models.Book
	if result := h.DB.First(&existingBook, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(fmt.Sprintf("Book with ID %d not found", id))
	} else {
		existingBook.Title = updatedBook.Title
		existingBook.Author = updatedBook.Author
		existingBook.Desc = updatedBook.Desc

		h.DB.Save(&existingBook)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(existingBook)
	}
}
