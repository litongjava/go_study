package main

import "net/http"

func test02010301() {
  //nil表示使用默认的DefaultServeMux
  http.ListenAndServe("localhost:8080", nil)
}

func test02010302() {
  server := http.Server{
    Addr:    "localhost:8080",
    Handler: nil,
  }
  server.ListenAndServe()
}

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello world"))
}
func test02010305() {
  mh := myHandler{}

  server := http.Server{
    Addr:    "localhost:8080",
    Handler: &mh,
  }
  server.ListenAndServe()
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

func (h *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("About"))
}
func test020108() {
  varHelloHandler := helloHandler{}
  varAboutHandler := aboutHandler{}

  server := http.Server{
    Addr:    "localhost:8080",
    Handler: nil, //DefaultServeMux
  }
  http.Handle("/hello", &varHelloHandler)
  http.Handle("/about", &varAboutHandler)

  server.ListenAndServe()
}
func welcome(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Welcome"))
}
func test020110() {
  server := http.Server{
    Addr:    "localhost:8080",
    Handler: nil, //DefaultServeMux
  }
  http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home"))
  })
  http.HandleFunc("/welcome", welcome)
  //將函数转为HandlerFunc
  http.HandleFunc("/welcome2", http.HandlerFunc(welcome))
  server.ListenAndServe()
}
func main() {
  test020110()
}
