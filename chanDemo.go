package main

import "fmt"

func main() {
	//var c chan int //声明，值nil
	//fmt.Printf("%T,%v\n", c, c)
	//c=make(chan int) // 创建
	//fmt.Printf("%T,%v\n", c, c)
	//a:=make(chan int) // 简短定义
	//fmt.Printf("%T,%v\n", a, a)

	a := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		a <- true
	}()
	if data, ok := <-a; ok {
		println(data)
		println("main quit")
	}

}
