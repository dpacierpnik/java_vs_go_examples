package endpoints

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Silesia Java Users!"))
	if err != nil {
		log.Println(err)
	}
}

