package main

import (
	"fmt"
	"time"
)

func main() {
	chann := make(chan int)
	go send(chann)

	// for range 访问通道
	for v := range chann {
		time.Sleep(1*time.Second)
		fmt.Println(v)
	}
}


func send(c chan int) {
	for i := 1; i <=50; i++ {
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}
		c <- i
	}
	close(c)
}
