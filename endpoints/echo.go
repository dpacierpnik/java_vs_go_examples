package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func EchoRequest(w http.ResponseWriter, r *http.Request) {

	data := &RequestData{
		QueryString: r.URL.RawQuery,
		Headers:     headersFrom(r),
	}

	body, readBodyErr := bodyFrom(r)
	if readBodyErr != nil {
		writeInternalServerError(w, "Marshalling data to JSON failed. Root cause: "+readBodyErr.Error())
		return
	}

	data.Body = body

	dataAsJson, marshallErr := json.MarshalIndent(data, "", " ")
	if marshallErr != nil {
		writeInternalServerError(w, "Marshalling data to JSON failed. Root cause: "+marshallErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, writeErr := w.Write(dataAsJson)
	if writeErr != nil {
		log.Println(writeErr)
		return
	}
}

func headersFrom(r *http.Request) []MyHeader {

	headers := make([]MyHeader, 0)
	for name, value := range r.Header {
		headers = append(headers, MyHeader{name, value})
	}
	return headers
}

func bodyFrom(r *http.Request) (string, error) {

	if r.Body == nil {
		return "", nil
	}

	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Println("WARNING: unable to close response body. Root cause: " + err.Error())
		}
	}()
	if readErr != nil {
		return "", readErr
	}
	return string(bodyBytes), nil
}

func writeInternalServerError(w http.ResponseWriter, errMsg string) {
	log.Println("ERROR:" + errMsg)
	w.Header().Set("Content-Type", "text/plain")
	http.Error(w, errMsg, http.StatusInternalServerError)
}

type RequestData struct {
	QueryString string     `json:"queryString,omitempty"`
	Headers     []MyHeader `json:"headers,omitempty"`
	Body        string     `json:"body,omitempty"`
}

type MyHeader struct {
	Name string `json:"name"`
	Value []string `json:"value"`
}
