package main

import "fmt"

func main() {
	// 匿名结构体
	s1:=struct {
		name string
		age int

	}{
		name:"张三",
		age:10,
	}
	fmt.Println(s1)

	s:=new(worker)
	s.string="sss"
	s.int=10
	fmt.Println(*s)
}

type worker struct {
	// 结构体匿名字段,默认使用类型做key
	string
	int
}