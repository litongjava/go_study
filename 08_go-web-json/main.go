package main

import (
  "08.go-web-json/controller"
  "08.go-web-json/middleware"
  "08.go-web-json/model"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  _ "net/http/pprof"
)

func main() {
  webEchoJson()
}

func webEchoJson() {
  controller.RegisterCompanyRoutes()
  //将TimeoutMiddleware的next设置为AuthMiddleware
  handler := &middleware.TimeoutMiddleware{
    Next: new(middleware.AuthMiddleware),
  }
  //导入pprof之后会自动添加一些handle,提供性能分析的数据接口,但是这些性能分析即可和自己编写的中间件并不怎么搭配,
  //所以通常再单独使用go runtine 单独监听一个节点,在这个频道上没有任何的中间件
  go http.ListenAndServe("localhost:8000", handler)
  http.ListenAndServe("localhost:8080", handler)
}

func marshalJson() {
  jsonStr := `{"id":123,"name":"Google","country":"USA"}`
  c := model.Company{}
  _ = json.Unmarshal([]byte(jsonStr), &c)
  fmt.Println(c)

  bytes, _ := json.Marshal(c)
  fmt.Println(string(bytes))

  bytes1, _ := json.MarshalIndent(c, "", "  ")
  fmt.Println(string(bytes1))
}

func webJson() {
  http.HandleFunc("/companies", homeHandle())
  http.ListenAndServe("localhost:8080", nil)
}

func homeHandle() func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
      //encode string to model
      dec := json.NewDecoder(r.Body)
      company := model.Company{}
      err := dec.Decode(&company)
      if err != nil {
        log.Println(err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
      }
      //decode model to string
      enc := json.NewEncoder(w)
      err = enc.Encode(company)
      if err != nil {
        log.Println(err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
      }
    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
    }
  }
}
