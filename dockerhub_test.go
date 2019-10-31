package dockerhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// assertBody asserts that the body of a given request is equal
// to the given string.
func assertBody(t *testing.T, req *http.Request, want string) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); want != got {
		fmt.Printf("got `%s`\n\n", got)
		fmt.Printf("want `%s`\n\n", want)
		t.Errorf("req.Body is %s; want %s", got, want)
	}
}

// hasHeader checks if a given header is present in an http.Request.
func hasHeader(req *http.Request, name string) bool {
	_, ok := req.Header[name]
	return ok
}

// assertHeader asserts that a given header is present and equal
// to the given value.
func assertHeader(t *testing.T, req *http.Request, name, want string) {
	if !hasHeader(req, name) {
		t.Errorf(`No header "%s" present`, name)
	}
	if got := req.Header.Get(name); got != want {
		t.Errorf(`req.Header["%s"] is %s; want %s`, name, got, want)
	}
}

// assertNoHeader asserts that a given header is not present in the
// http.Request.
func assertNoHeader(t *testing.T, req *http.Request, name string) {
	if hasHeader(req, name) {
		t.Errorf(`Expected header "%s" not to be present`, name)
	}
}

// assertMethod asserts that an http.Request is of a given method.
func assertMethod(t *testing.T, req *http.Request, want string) {
	if got := req.Method; got != want {
		t.Errorf("req.Method is %s; want %s", got, want)
	}
}

// assertMarshalledJSON asserts that a value is marshalled as intended.
func assertMarshalledJSON(t *testing.T, v interface{}, want string) {
	b, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Error marshalling value %v: %v", v, err)
	}
	if got := string(b); got != want {
		t.Errorf("marshalled to %s; want %s", got, want)
	}
}

// makeMockClient spins up a local server with mux and returns a Client
// pointing to it. A handler should be added to mock the endpoint being
// targeted.
func makeMockClient() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	handler := http.NewServeMux()
	handler.Handle(defaultAPIBaseEndpoint+"/", http.StripPrefix(defaultAPIBaseEndpoint, mux))
	srv := httptest.NewServer(handler)

	client = NewClient(nil)
	url, _ := url.Parse(srv.URL + defaultAPIBaseEndpoint + "/")
	client.BaseURL = url

	return client, mux, srv.Close
}

// mustJSONMarshal marshals a value to JSON or panics.
func mustJSONMarshal(v interface{}) []byte {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
