package main

import (
	"com.jamf.services.java_vs_go/clients"
	"com.jamf.services.java_vs_go/endpoints"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	httpbinClient := &clients.Httpbin{
		Url: "http://localhost/json",
	}
	httpbinEndpoint := endpoints.HttpbinEndpoint{
		Client: httpbinClient,
	}

	r := mux.NewRouter()
	r.Path("/hello").Methods(http.MethodGet).HandlerFunc(endpoints.Hello)
	r.Path("/echo-request").Methods(http.MethodPost, http.MethodGet).HandlerFunc(endpoints.EchoRequest)
	r.Path("/httpbin/json").Methods(http.MethodGet).HandlerFunc(httpbinEndpoint.RewriteJson)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
