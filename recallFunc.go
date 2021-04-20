package main

import "fmt"

func main() {
	s:=operate(1, 2, add)
	println(s)

	sub:=func(a,b int)int{
		return a-b
	}

	s1:=operate(2, 2, sub)
	println(s1)
}

func add(a, b int) int {
	return a + b
}

func operate(a int, b int, f func(int, int) int) int {
	fmt.Println("operator pre")
	sum := f(a, b)
	fmt.Printf("calc:%d\n",sum)
	fmt.Println("operator after")
	return sum
}
