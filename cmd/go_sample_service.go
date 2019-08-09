package main

import (
	"com.jamf.services.java_vs_go/endpoints"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", endpoints.Hello)
	http.HandleFunc("/echo-request", endpoints.EchoRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
