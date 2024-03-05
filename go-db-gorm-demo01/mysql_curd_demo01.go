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
  db.Create(&model.Product{Code: "001", Price: 100})

  var product model.Product
  // find product with integer primary key
  db.First(&product, 1)
  // find product with code D42
  db.First(&product, "code=?", "D42")

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.Model(&product).Updates(model.Product{Price: 200, Code: "F42"})
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  //delete
  db.Delete(&product, 1)

}
