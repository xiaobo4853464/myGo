package main

import "fmt"

func main() {
	go printNum()
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("main quit")

}

func printNum() {
	for i := 0; i < 100; i++ {
		fmt.Printf("sub print:%d\n", i)
	}
}
