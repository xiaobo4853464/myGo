package main

import "fmt"

func main() {
	const (
		// iota 会累计加1
		A=iota
		B=iota
		C=iota
		D=iota
	)
	const (
		// iota 会累计加1，如果常量为赋值，与上行一样
		A1=iota
		B2
		C3
		D4
	)
	fmt.Println(A,B,C,D)
	fmt.Println(A1,B2,C3,D4)
}
