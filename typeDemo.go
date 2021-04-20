package main

import "fmt"

func main() {
	var a myint
	var b mystr = "hello"
	var a1 int = 11
	var b1 string = "hello1"
	fmt.Println(a, b)
	fmt.Println(a1, b1)

	//a = a1 //cannot use a1 (type int) as type myint in assignment
	//b = b1 //cannot use b1 (type string) as type mystr in assignment

	fmt.Println("=======================")
	var aa myint1
	var aa1 int
	fmt.Printf("%T,%T",aa,aa1)

}

// 新建一个自己的类型，不想等
type mystr string
type myint int
//取别名，相等
type myint1 = int
