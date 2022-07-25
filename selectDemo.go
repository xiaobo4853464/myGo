package main

import (
	"fmt"
	"time"
)

func main() {
	/*select 类似于switch，用于channel
	select 会随机的执行一个可运行的case
	如果没有select可以执行，则执行default语句，否则进入阻塞， 直到某个case语句可运行*/
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中的数据", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中的数据", num2)
		} else {
			fmt.Println("ch2已经关闭")
		}
	default:
		fmt.Println("执行default")
	}
}
