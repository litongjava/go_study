package model

import (
  "database/sql"
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type User struct {
  gorm.Model
  MemberNumber sql.NullString
  Price        uint
}

type Student struct {
  gorm.Model
  Name   string
  Age    uint
  Gender uint
}

//func (User) TableName() string {
//  return "my_user"
//}
