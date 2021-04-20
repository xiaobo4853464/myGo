package main

import "fmt"

func main() {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg, "恢复")
		}
	}()
	panicF()
}

func panicF() {
	defer fmt.Println("11111")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			panic("panic!!!") // panic 后的所有都不执行
		}
	}
	defer fmt.Println("22222")
}
