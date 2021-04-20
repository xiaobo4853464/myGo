package main

import (
	"fmt"
	"time"
)

func main() {
	ccc := make(chan int)

	go func() {
		time.Sleep(1*time.Second)
		data := <-ccc
		fmt.Println(data)
	}()

	time.Sleep(3*time.Second)
	fmt.Println("开始写入")
	ccc <- 100
}
