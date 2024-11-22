package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestUnpack(t *testing.T) {
	type reqData struct {
		Name string   `mytag:"n"`
		Age  int      `mytag:"a"`
		Jobs []string `mytag:"j"`
	}
	expected := reqData{
		Name: "Nikita",
		Age:  25,
		Jobs: []string{"developer", "gopher"},
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080?n=Nikita&a=25&j=developer&j=gopher", nil)
	got := reqData{}

	err := Unpack(req, &got)
	if err != nil {
		t.Fatal("unpack error", err)
	}
	if !reflect.DeepEqual(got, expected) {
		gotBytes, _ := json.Marshal(got)
		expectedBytes, _ := json.Marshal(expected)
		t.Log("got: ", string(gotBytes))
		t.Log("expected: ", string(expectedBytes))
		t.Fatal("got and expected not equal")
	}
}
