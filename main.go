package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", ExampleHandler)

	log.Println("Service Started on Port 443")
	err := http.ListenAndServe(":443", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok", "msg":"railway-go-http-001_2023-07-05"}`)
}
