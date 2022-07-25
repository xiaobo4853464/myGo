package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int32 = 10
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
}
