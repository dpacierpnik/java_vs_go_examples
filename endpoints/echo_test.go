package endpoints


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldEchoQueryString(t *testing.T) {

	// GIVEN:
	req, err := http.NewRequest(http.MethodGet, "?param1=value1&param2=value2", nil)
	if err != nil {
		t.Fatal(err)
	}

	respRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(EchoRequest)

	// WHEN:
	handler.ServeHTTP(respRecorder, req)

	// THEN:
	if statusCode := respRecorder.Code; statusCode != http.StatusOK {
		t.Errorf("handler returned wrong statusCode code: got %v want %v", statusCode, http.StatusOK)
	}

	if contentType := respRecorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", contentType, "application/json")
	}

	expectedResponseBody := `{
 "queryString": "param1=value1\u0026param2=value2"
}`
	if respRecorder.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", respRecorder.Body.String(), expectedResponseBody)
	}
}

func TestShouldEchoHeaders(t *testing.T) {

	// GIVEN:
	req, err := http.NewRequest(http.MethodGet, "/echo-request", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Tenant-Id", "some-tenant-id")

	respRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(EchoRequest)

	// WHEN:
	handler.ServeHTTP(respRecorder, req)

	// THEN:
	if statusCode := respRecorder.Code; statusCode != http.StatusOK {
		t.Errorf("handler returned wrong statusCode code: got %v want %v", statusCode, http.StatusOK)
	}

	if contentType := respRecorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", contentType, "application/json")
	}

	expectedResponseBody := `{
 "headers": [
  {
   "name": "Tenant-Id",
   "value": [
    "some-tenant-id"
   ]
  }
 ]
}`

	if respRecorder.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", respRecorder.Body.String(), expectedResponseBody)
	}
}

func TestShouldEchoBody(t *testing.T) {

	// GIVEN:
	bodyReader := bytes.NewBufferString(`{"message": "Hello Silesia Java Users!"}`)
	req, err := http.NewRequest(http.MethodGet, "/echo-request", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	respRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(EchoRequest)

	// WHEN:
	handler.ServeHTTP(respRecorder, req)

	// THEN:
	if statusCode := respRecorder.Code; statusCode != http.StatusOK {
		t.Errorf("handler returned wrong statusCode code: got %v want %v", statusCode, http.StatusOK)
	}

	if contentType := respRecorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", contentType, "application/json")
	}

	expectedResponseBody := `{
 "body": "{\"message\": \"Hello Silesia Java Users!\"}"
}`

	if respRecorder.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", respRecorder.Body.String(), expectedResponseBody)
	}
}
