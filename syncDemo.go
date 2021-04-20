package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)

	go f1()
	go f2()

	wg.Wait()
	fmt.Println("main quit")
}

func f1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("func1 ", i)
	}

}

func f2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("func2 ", i)
	}
}
