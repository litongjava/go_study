package controller

import (
  "08.go-web-json/model"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandleCompnayCorrect(t *testing.T) {
  request := httptest.NewRequest(http.MethodGet, "/companies", nil)
  writer := httptest.NewRecorder()

  handleCompnay(writer, request)
  result, _ := ioutil.ReadAll(writer.Result().Body)

  company := model.Company{}
  json.Unmarshal(result, &company)

  if company.ID != 123 {
    t.Errorf("Failed")
  }
}
