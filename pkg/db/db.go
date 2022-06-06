package db

import (
  "go-rest-api-beginner-karan/pkg/models"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)

func Init() *gorm.DB {
  dbUrl := "postgres://pg:pass@localhost:5432/crud"

  db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

  if err != nil {
    log.Fatalln(err)
  }

  db.AutoMigrate(&models.Book{})

  return db
}
