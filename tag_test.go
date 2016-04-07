package godt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestListTags(t *testing.T) {
	setUp()
	defer tearDown()

	type Tag struct {
		Name string
	}
	image := "foo"
	path := fmt.Sprintf(tagsPath[client.APIVersion], image)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")

		tag := Tag{"bar"}
		res, _ := json.Marshal(tag)

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	resp, err := client.ListTags(image)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ListTags(): %v", err)
	}

	var foo Tag

	_ = json.Unmarshal(body, &foo)
	expected := Tag{"bar"}
	if !reflect.DeepEqual(foo, expected) {
		t.Errorf("Expected %v - Got %v", expected, foo)
	}
}
