package main

import "fmt"

func main() {
	var arr [4]int
	//arr[0] = 1
	//println(arr[0])
	//println(arr[1])
	//println(arr[2])
	//println(arr[3])
	for _,v := range arr {
		fmt.Println(v)
	}

	//  数组的长度和容量是一样的
	fmt.Printf("数组的长度:%d\n", len(arr))
	fmt.Printf("数组的容量:%d\n", cap(arr))

}
