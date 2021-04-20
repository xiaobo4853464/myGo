package main

import (
	"fmt"
	"time"
)

func main() {
	// 单向通道只用来参数定义
	c := make(chan int)
	go sssend(c)
	go recv(c)
	time.Sleep(3 * time.Second)
}

func sssend(c chan<- int) { //只写
	c <- 100
	fmt.Println("sent")

}

func recv(c <-chan int) { //只读
	data := <-c
	fmt.Println("recv:", data)

}
