package main

import "fmt"

func main() {
	var a int8=100
	var b int16
	var c float64 =4.32
	b=int16(a)
	fmt.Println(b)
	b=int16(c)
	fmt.Println(b)
}
