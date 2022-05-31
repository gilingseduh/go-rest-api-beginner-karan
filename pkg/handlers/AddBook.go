package handlers

import (
	"encoding/json"
	"go-rest-api-beginner-karan/pkg/mocks"
	"go-rest-api-beginner-karan/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book

	json.Unmarshal(body, &book)

	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
