package main

import "fmt"

func main() {
	var a = 10
	var p *int // int 指针
	fmt.Println(p)
	p = &a          //赋值要赋地址
	fmt.Println(&a) // 打印a的地址
	fmt.Println(p)  // 打印a的地址
	fmt.Println(&p) // 打印p的地址，指针的地址
	fmt.Println(*p) // 打印值，也就是a的值

	a = 11

	fmt.Println(a)
	fmt.Println(*p)

	fmt.Println("----------------")

	b := 1
	var p2 *int
	p2 = &b
	bb := pFunc(p2)
	fmt.Println(bb)

	//var arr1= [4]int{1,2,3,4}
	//
	//var p1 *[4] int
	//
	//p1=&arr1
	//fmt.Println(&arr1)
	//fmt.Println(p1)
	//fmt.Println(*p1)
	//
	//p1[0]=100
	//fmt.Println(arr1)
	//fmt.Println(*p1)

}

func pFunc(p *int) int {
	fmt.Println(p)
	s := *p + 1
	fmt.Println(s)
	return s
}
