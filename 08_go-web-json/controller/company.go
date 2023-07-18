package controller

import (
  "08.go-web-json/model"
  "encoding/json"
  "net/http"
  "time"
)

func RegisterCompanyRoutes() {
  http.HandleFunc("/companies", handleCompnay)
}

func handleCompnay(writer http.ResponseWriter, request *http.Request) {
  c := model.Company{
    ID:      123,
    Name:    "Google",
    Country: "USA",
  }
  time.Sleep(4 * time.Second)
  enc := json.NewEncoder(writer)
  enc.Encode(c)
}
