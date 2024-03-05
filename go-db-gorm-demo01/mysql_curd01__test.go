package main

import (
  "go-db-gorm-demo01/model"
  "testing"
)

func TestYourFunction(t *testing.T) {
  db, err := open()
  if err != nil {
    panic(err)
  }
  db.Create(&model.Product{
    Code:  "001",
    Price: 100,
  })

  print("crate")

}
