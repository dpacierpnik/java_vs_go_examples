package endpoints

import (
	"log"
	"net/http"
)

type HttpbinClient interface {
	Json() (string, error)
}

type HttpbinEndpoint struct {
	Client HttpbinClient
}

func (e *HttpbinEndpoint) RewriteJson(w http.ResponseWriter, r *http.Request) {

	jsonFromHttpBin, clientErr := e.Client.Json()
	if clientErr != nil {
		log.Println(clientErr)
		http.Error(w, "Httpbin client error: "+clientErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, writeErr := w.Write([]byte(jsonFromHttpBin))
	if writeErr != nil {
		log.Println(writeErr)
		http.Error(w, "Error writing response. Root cause: "+writeErr.Error(), http.StatusInternalServerError)
		return
	}
}
