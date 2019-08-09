package main

import (
	"log"
	"net/http"
)

func main() {
	// registers `hello` handler function under `/hello` path in the default HTTP router
	http.HandleFunc("/hello", hello)
	// runs HTTP service on 8080 port on localhost
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Silesia Java Users!"))
	if err != nil {
		log.Println(err)
	}
}
