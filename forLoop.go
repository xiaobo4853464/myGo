package main

import "fmt"

func main() {
	// 相当于while true
	//i:=0
	//for{
	//	fmt.Println(i)
	//	i++
	//}

	a:=0
	label:
		fmt.Println("label")
	for a<=10 {
		if a==8{
			a++
			goto label
		}
		fmt.Println(a)
		a++
	}
}
