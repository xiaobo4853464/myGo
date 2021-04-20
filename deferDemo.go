package main

import "fmt"

func main() {
	//multiDefer()

	a:=1
	fmt.Println(a)
	defer func1(a) //调用defer时候，参数已经传递了，只是延时执行
	a++
	fmt.Println(a)
}


// 多个defer，从下往上
func multiDefer(){
	defer fmt.Println("hello")
	fmt.Println("-----")
	defer fmt.Println("world")
	fmt.Println("-----")
}

func func1(a int){
	fmt.Printf("调用中的：%d",a)
}