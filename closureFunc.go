package main

import "fmt"

func main() {
	fmt.Println(fc1())
	fmt.Println(fc1())
}

func fc1() int {
	i := 0
	a := func(i int) int {
		i++
		return i
	}(i)
	return a
}
