package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("server model")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("message")
		msg := "Hello World !!!"
		if len(param) > 0 {
			msg = param
		}
		fmt.Fprintf(w, "<h1> %v </h1>", msg)
	})
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
