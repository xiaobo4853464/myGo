package main

import "fmt"

func main() {
	// 继承
	h:=Human{"human",10}
	man1:=Man{Human:h,gender: "male"}
	woman1:=Woman{Human:h,gender: "female"}
	man1.eat()
	woman1.eat()


}

//super
type Human struct {
	name string
	age  int
}

type Man struct {
	Human
	gender string
}

type Woman struct {
	Human
	gender string
}

func (h *Human) eat()  {
	fmt.Println("super's method, eating")
}

func (m *Man) eat()  {
	fmt.Println("child [man]'s method, eating")
}