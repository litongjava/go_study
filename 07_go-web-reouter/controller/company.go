package controller

import (
  "fmt"
  "net/http"
  "regexp"
  "strconv"
)

func registerCompanyRouters() {
  http.HandleFunc("/companies", handleCompanies)
  http.HandleFunc("/companies/", handleCompany)
}

func handleCompanies(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintln(writer, "Google", "MicroSoft", "Meta")
}
func handleCompany(writer http.ResponseWriter, request *http.Request) {
  pattern, _ := regexp.Compile(`/companies/(\d+)`)
  matches := pattern.FindStringSubmatch(request.URL.Path)
  if len(matches) > 0 {
    for i := 0; i < len(matches); i++ {
      fmt.Println(matches[i])
    }
    //将string转为int
    companyId, _ := strconv.Atoi(matches[1])
    fmt.Fprintln(writer, companyId)
  } else {
    writer.WriteHeader(http.StatusNotFound)
  }
}
