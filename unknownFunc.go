package main

import "fmt"

func main() {
	func() {
		println("匿名函数，调用一次")
	}()

	f := func() {
		println("匿名函数，可以被调用N次")
	}
	f()
	f()

	// 有参数，带返回的匿名函数
	r := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(r)
}
