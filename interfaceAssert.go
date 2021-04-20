package main

import (
	"fmt"
	"math"
)

func main() {
	t := Triangle{1, 2, 3}
	c := Circle{4}
	getType(&t)
	getType(&c)
}

func getType(s Shape) {
	// 接口断言，不会panic
	if obj, ok := s.(*Triangle); ok {
		fmt.Printf("三角形：%.2f,%.2f,%.2f\n", obj.a, obj.b, obj.c)
	} else if obj, ok := s.(*Circle); ok {
		fmt.Printf("圆：%.2f\n", obj.radius)
	}
}

type Shape interface {
	peri() float64
	area() float64
}

type Triangle struct {
	a, b, c float64
}

type Circle struct {
	radius float64
}

func (t *Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t *Triangle) area() float64 {
	p := t.peri() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c *Circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
