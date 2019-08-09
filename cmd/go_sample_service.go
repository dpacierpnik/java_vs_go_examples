package main

import (
	"com.jamf.services.java_vs_go/clients"
	"com.jamf.services.java_vs_go/endpoints"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", endpoints.Hello)
	http.HandleFunc("/echo-request", endpoints.EchoRequest)

	httpbinClient := &clients.Httpbin{
		Url: "http://localhost/json",
	}
	httpbinEndpoint := endpoints.HttpbinEndpoint{
		Client: httpbinClient,
	}
	http.HandleFunc("/httpbin/json", httpbinEndpoint.RewriteJson)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
