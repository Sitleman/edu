package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

func Unpack(req *http.Request, target any) error {
	params, _ := url.ParseQuery(req.URL.RawQuery)
	paramsBytes, _ := json.Marshal(params)
	fmt.Println("params: ", string(paramsBytes))

	v := reflect.ValueOf(&target)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target must be a pointer to struct, v.Kind() = %s", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("mytag")
		if tag == "" {
			tag = field.Name
		}
		fmt.Println("field: ", field)

		values, ok := params[tag]
		if !ok {
			continue
		}
		fmt.Println("value: ", values)

		vv := v.Field(i)
		switch v.Kind() {
		case reflect.String:
			vv.SetString(values[0])
		case reflect.Slice:
			vv.Set(reflect.ValueOf(values))
		case reflect.Int:
			vNum, _ := strconv.Atoi(values[0])
			vv.SetInt(int64(vNum))
		}
		//field.Type
		//v.Field(i).SetString()
	}
	fmt.Println("v:", v)
	return nil
}
