package main

import "fmt"

func main() {
	/*
		接口对象不能访问实现对象中的属性
	*/
	m := Mouse{name: "logic"}
	f := FlashDisk{name: "samsung"}
	testInterface(&m)
	testInterface(&f)

	// 结构嵌套
	fmt.Println("------结构嵌套---------")
	d := new(Dog)
	var c C
	c = d
	c.test1()
	c.test2()
	c.test3()

}

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

func (m *Mouse) start() {
	fmt.Printf("%s click click\n", m.name)
}

func (m *Mouse) end() {
	fmt.Printf("%s mouse quit\n", m.name)
}

func (f *FlashDisk) start() {
	fmt.Printf("%s read or write data\n", f.name)

}

func (f *FlashDisk) end() {
	fmt.Printf("%s pop\n", f.name)
}

func testInterface(u USB) {
	u.start()
	u.end()
}

type A interface {
	test1()
}
type B interface {
	test2()
}
type C interface {
	A
	B
	test3()
}

type Dog struct {
}

func (d *Dog) test1() {
	fmt.Println("test1...")
}
func (d *Dog) test2() {
	fmt.Println("test2...")
}
func (d *Dog) test3() {
	fmt.Println("test3...")
}
