package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	s2 := make([]int, 3,5)
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3

	//s = append(s, 1,2,3)
	s = append(s, s2...)
	for _, i2 := range s {
		fmt.Println(i2)
	}

	a:=[4]int{1,2,3,4}

	// 从数组创建slice
	s3:=a[:]
	fmt.Println(s3)
	// 修改底层数组会直接修改切片
	a[0]=0
	fmt.Println(s3)


}
