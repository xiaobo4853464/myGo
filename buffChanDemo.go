package main

import "fmt"



func main() {
	c := make(chan int, 5) //带缓冲的channel
	fmt.Println(len(c), cap(c))

	go ssend(c)
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("main quit...")
}

func ssend(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("写数据：", i)
		c <- i
	}
	close(c)
}
