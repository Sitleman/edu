package main

import (
	"fmt"
	"reflect"
)

func Display(val interface{}) {
	display("root", val)
}

func display(path string, val interface{}) {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	fmt.Printf("%s: %v (%s) \n", path, v, t)
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			display(fmt.Sprintf("%s.%s", path, t.Field(i).Name), v.Field(i).Interface())
		}
	}
}
