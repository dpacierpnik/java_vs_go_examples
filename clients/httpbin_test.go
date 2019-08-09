package clients

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpbinJson(t *testing.T) {

	// GIVEN:
	// you can use also url.Values to build this
	payload := `{
     "slideshow": {
       "author": "Yours Truly",
       "date": "date of publication",
       "slides": [
         {
           "title": "Wake up to WonderWidgets!",
           "type": "all"
         },
         {
           "items": [
             "Why <em>WonderWidgets</em> are great",
             "Who <em>buys</em> WonderWidgets"
           ],
           "title": "Overview",
           "type": "all"
         }
       ],
       "title": "Sample Slide Show"
     }
   }`

	var httpbinStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		if _, writeErr := w.Write([]byte(payload)); writeErr != nil {
			t.Fatalf("Error writing stubbed data. Root cause: %v", writeErr)
		}
	}))

	clientUT := Httpbin{
		Url: httpbinStub.URL,
	}

	// WHEN:
	actualJson, err := clientUT.Json()

	// THEN:
	if err != nil {
		t.Fatal("Error should be nil")
	}

	if actualJson == "" {
		t.Errorf("handler returned unexpected body: got %v want %v", actualJson, payload)
	}
}
