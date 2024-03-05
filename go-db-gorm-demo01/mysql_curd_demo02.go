package main

import (
  "database/sql"
  "go-db-gorm-demo01/model"
  "go-db-gorm-demo01/utils"
)

func main() {
  db, err := utils.Open()
  if err != nil {
    panic(err)
  }
  db.AutoMigrate(&model.User{})
  var user model.User

  nullStringFor20 := sql.NullString{
    String: "20",
    Valid:  true,
  }

  db.Create(&model.User{Price: 100, MemberNumber: nullStringFor20})

  nullString := sql.NullString{
    String: "",
    Valid:  true,
  }
  db.Model(&user).Where("id = 1").Updates(model.User{Price: 100, MemberNumber: nullString}) // non-zero fields
}
