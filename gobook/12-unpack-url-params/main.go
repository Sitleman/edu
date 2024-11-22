package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MyType uint64
	var aa uint64 = 1
	var a = MyType(aa)
	fmt.Println("a: ", a)
	fmt.Println("reflect.TypeOf(a): ", reflect.TypeOf(a))
	fmt.Println("reflect.TypeOf(a).Kind(): ", reflect.TypeOf(a).Kind())
	fmt.Println("reflect.ValueOf(a): ", reflect.ValueOf(a))
	fmt.Println("reflect.ValueOf(a).Kind(): ", reflect.ValueOf(a).Kind())
	fmt.Println("reflect.ValueOf(a).Type(): ", reflect.ValueOf(a).Type())
	fmt.Println("reflect.ValueOf(a).CanSet(): ", reflect.ValueOf(a).CanSet())
	fmt.Println("reflect.ValueOf(a).CanAddr(): ", reflect.ValueOf(a).CanAddr())
	fmt.Println("reflect.ValueOf(&a).Elem(): ", reflect.ValueOf(&a).Elem())

	var b interface{} = &aa
	fmt.Println("b: ", b)
	fmt.Println("reflect.ValueOf(b).CanSet(): ", reflect.ValueOf(b).CanSet())
	fmt.Println("reflect.ValueOf(b).CanSet(): ", reflect.ValueOf(b).CanAddr())

	var c uint64 = 1
	cc := reflect.ValueOf(&c).Elem()
	cc.SetUint(5)
	fmt.Println(c)
	//
	//var d uint64 = 1
	//dd := reflect.ValueOf(d)
	//dd.SetUint(5)
	//fmt.Println(d)

	//fmt.Println(a)
}
