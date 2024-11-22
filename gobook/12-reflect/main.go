package main

import "encoding/json"

func main() {
	val1 := 42
	val2 := struct {
		X    int
		Name string
	}{42, "hehehe"}
	var val3 interface{}
	json.Unmarshal([]byte(`{"X": 42, "Name": "hehehe"}`), &val3)
	_, _, _ = val1, val2, val3
	Display(val2)
}
