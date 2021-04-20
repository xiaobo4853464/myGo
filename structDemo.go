package main

import "fmt"

func main() {
	var p Person
	p.name = "金迪"
	p.age = 25
	p.gender = "女"

	fmt.Println(p)

	p1 := Person{name: "shawn", age: 28, gender: "male"}
	fmt.Println(p1)

	fmt.Println("-------指针访问结构体-------")

	// new 专门构建某种类型的指针函数，不是nil，初始化为0值
	p2 := new(Person)
	var p3 *Person // 空指针
	fmt.Println(p2 == nil)
	fmt.Println(p3 == nil)

	fmt.Println(p2)
	fmt.Println(p3)
	p2.name = "aba"
	p2.age = 32
	p2.gender = "女"
	fmt.Println(*p2)

	fmt.Println("-------嵌套结构体-------")
	b1 := new(book_)
	b1.name = "java"
	b1.price = 59.1

	s1 := new(student)
	s1.name = "xiaobo"
	s1.book = *b1
	fmt.Println(*s1)
	fmt.Println(s1.book)

}

type Person struct {
	name   string
	age    int
	gender string
}

// 结构体嵌套
type book_ struct {
	name  string
	price float64
}

type student struct {
	name string
	book book_
}
