package godt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

func setUp() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewHTTPClient(nil)
	url, _ := url.Parse(server.URL)
	client.HubURL = url
}

func tearDown() {
	server.Close()
}

func assertEqual(t *testing.T, result interface{}, expect interface{}) {
	if result != expect {
		t.Errorf("Expect (Value: %v) (Type: %T) - Got (Value: %v) (Type: %T)", expect, expect, result, result)
	}
}

func TestNewHTTPClient(t *testing.T) {
	c := NewHTTPClient(nil)

	assertEqual(t, c.UserAgent, ua)
	assertEqual(t, c.HubURL.String(), defaultHubURL)
	assertEqual(t, c.APIVersion, defaultAPIVersion)
}

func TestNewrequest(t *testing.T) {
	c := NewHTTPClient(nil)

	req, _ := c.NewRequest("GET", "/foo", nil)
	assertEqual(t, req.URL.String(), defaultHubURL+"/foo")
	assertEqual(t, req.Header.Get("User-Agent"), ua)
	assertEqual(t, req.Header.Get("Content-Type"), mediaType)
	assertEqual(t, req.Header.Get("Accept"), mediaType)
}

func TestDo(t *testing.T) {
	setUp()
	defer tearDown()

	type Foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")

		foo := Foo{"bar"}
		res, _ := json.Marshal(foo)

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	req, _ := client.NewRequest("GET", "/", nil)

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	var foo Foo

	_ = json.Unmarshal(body, &foo)
	expected := Foo{"bar"}
	if !reflect.DeepEqual(foo, expected) {
		t.Errorf("Expected %v - Got %v", expected, foo)
	}
}
