package main

import "fmt"

func main() {
	w1 := Worker{name:"肖肖"}
	c1 := Cat{"小白"}
	w1.sleep()
	c1.sleep()
	//var w2 =Worker{name:"aaa",age:1}
	//w2.sleep()
}

type Worker struct {
	name string
}

type Cat struct {
	name string
}

func (w *Worker) sleep() {
	fmt.Printf("%s sleeping\n", w.name)
}

func (c *Cat) sleep() {
	fmt.Printf("%s sleeping\n", c.name)
}