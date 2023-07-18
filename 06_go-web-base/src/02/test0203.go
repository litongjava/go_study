package main

import (
	"fmt"
	"log"
	"net/http"
)

func test020307() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.Fragment)
	})

	server.ListenAndServe()
}
func test020309() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Header)
		fmt.Fprintln(w)
		fmt.Fprintln(w, r.Header["Accept-Encoding"])
		fmt.Fprintln(w)
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
	})

	server.ListenAndServe()
}
func test020310() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	server.ListenAndServe()
}
func test020311() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query() //map[string][]string

		id := query["id"] //[]string
		log.Println(id)
		name := query.Get("name") //string
		log.Println(name)
	})

	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	test020311()
}
