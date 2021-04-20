package main

import "fmt"

func main() {
	var a =10
	var p *int
	fmt.Println(p)
	p=&a
	fmt.Println(&a) // a 的地址
	fmt.Println(p) // p 的值 是 a 的地址
	fmt.Println(*p) // p 的数字， 就是a 的值
	fmt.Println(&p) // p 的地址

	a=11

	fmt.Println(a)
	fmt.Println(*p)

	fmt.Println("----------------")

	b:=1
	var p2 *int
	p2=&b
	bb:=pFunc(p2)
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

func pFunc(p *int)int{
	fmt.Println(p)
	s:=*p+1
	fmt.Println(s)
	return s
}
