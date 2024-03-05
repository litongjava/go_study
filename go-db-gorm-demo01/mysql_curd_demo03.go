package main

import (
  "go-db-gorm-demo01/model"
  "go-db-gorm-demo01/utils"
)

func main() {
  db, err := utils.Open()
  if err != nil {
    panic(err)
  }
  db.AutoMigrate(&model.Student{})

  value := model.Student{
    Name:   "Tong Li",
    Age:    18,
    Gender: 1,
  }
  db.Create(&value)

  var students = []model.Student{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
  db.CreateInBatches(students, 100)
  //db.Create(students)

}
