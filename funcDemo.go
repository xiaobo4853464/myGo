package main

import "fmt"

func main() {
	a := getSum(1, 5)
	fmt.Println(a)

	s := []int{100, 200}
	aa := getSumByArgs(s...)
	fmt.Println(aa)

	x,y:=getMultiReturn(1,1)
	fmt.Println(x,y)

}

//函数参数一致，可以简写在一起
//func getSum(start, end int) int {
func getSum(start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}

	return sum

}

//可变参数
func getSumByArgs(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum

}

// 多返回
func getMultiReturn(a, b int) (int, int) {
	return a + 1, b + 1
}
