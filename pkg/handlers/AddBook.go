package handlers

import (
  "encoding/json"
  "fmt"
  "go-rest-api-beginner-karan/pkg/models"
  "io/ioutil"
  "log"
  "net/http"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    log.Fatalln(err)
  }

  var book models.Book

  json.Unmarshal(body, &book)

  if result := h.DB.Create(&book); result.Error != nil {
    fmt.Println(result.Error)
  }

  w.WriteHeader(http.StatusCreated)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("Created")
}
