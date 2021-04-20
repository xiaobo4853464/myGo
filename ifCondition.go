package main

import "fmt"

func main() {
	if num := 1; num%2 == 0 { // num只能在if语句中使用
		fmt.Printf("true %d",num)
	} else {
		fmt.Printf("false %d",num)
	}

}
