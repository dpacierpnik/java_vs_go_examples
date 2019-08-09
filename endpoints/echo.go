package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// public void echoRequest(ResponseWriter w, Request r)
func EchoRequest(w http.ResponseWriter, r *http.Request) {

	// var data = new RequestDate(r.URL.RawQuery, headersFrom(r));
	// OR
	// var dataBuilder = RequestDate.newBuilder()
	//   .withQueryString(r.URL.RawQuery)
	//   .withHeaders(headersFrom(r));
	data := &RequestData{
		QueryString: r.URL.RawQuery,
		Headers:     headersFrom(r),
	}

	// String body = null;
	// try {
	// 	 body = bodyFrom(r);
	// } catch(ReadBodyException e) {
	// 	 writeInternalServerError(w, "Marshalling data to JSON failed. Root cause: " + e.getMessage());
	// 	 return;
	// }
	body, readBodyErr := bodyFrom(r)
	if readBodyErr != nil {
		writeInternalServerError(w, "Marshalling data to JSON failed. Root cause: "+readBodyErr.Error())
		return
	}

	// data.setBody(body);
	// OR
	// dataBuilder.withBody(body);
	data.Body = body

	// String dataAsJson = null;
	// try {
	// 	 var objectWriter = new ObjectMapper().writerWithDefaultPrettyPrinter();
	//   var dataAsJson = objectWriter.writeValueAsString(data);
	// } catch (JsonProcessingException e) {
	//   writeInternalServerError(w, "Marshalling data to JSON failed. Root cause: " + e.getMessage());
	//   return;
	// }
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

// []MyHeader headersFrom(Request r)
func headersFrom(r *http.Request) []MyHeader {

	// var headers = new ArrayList<MyHeader>();
	headers := make([]MyHeader, 0)
	// for (var entry : headers1.entrySet()) {
	for name, value := range r.Header {
		// headers.add(new MyHeader(entry.getKey(), entry.getValue()));
		headers = append(headers, MyHeader{name, value})
	}
	return headers

	//return r.getHeaders().entrySet().stream()
	//  .map(entry -> new MyHeader(entry.getKey(), entry.getValue()))
	//  .collect(Collectors.toList());
}

// String bodyFrom(Request r) throws IOException {
func bodyFrom(r *http.Request) (string, error) {

	// if(r.getBody() == null) {
	if r.Body == nil {
		// return null;
		return "", nil
	}

	// try(var bodyReader = new BufferedReader(r.Body)) {
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	//defer func() {
	//	err := r.Body.Close()
	//	if err != nil {
	//		log.Println("WARNING: unable to close response body. Root cause: " + err.Error())
	//	}
	//}()
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

// @JsonInclude(NON_NULL)
// class RequestData {
//
// 	   @JsonProperty("queryString")
//     public String QueryString;
//
//     @JsonProperty("headers")
//     public Header[] Headers;
//
//     @JsonProperty("body")
//     public string Body;
//
// 	   string anyField;
//
type RequestData struct {
	QueryString string     `json:"queryString,omitempty"`
	Headers     []MyHeader `json:"headers,omitempty"`
	Body        string     `json:"body,omitempty"`
	anyField    string
}

// public String getAnyField() {
//     return anyField;
// }
// by convention Getter does not start with Get
func (r RequestData) AnyField() string {
	return r.anyField
}

// Setter
func (r RequestData) SetAnyField(v string) {
	r.anyField = v
}

// class MyHeader {
//
// 		@JsonProperty("name")
// 		public String name;
//
// 		@JsonProperty("value")
// 		public String[] value;
// }
type MyHeader struct {
	Name string `json:"name"`
	Value []string `json:"value"`
}
