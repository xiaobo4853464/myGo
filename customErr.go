package main

import (
	"fmt"
	"math"
)

func main() {
	x, e := circleArea(2)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(x)
}

// 自定义的错误
type areaErr struct {
	msg    string
	radius float64
}

func (e *areaErr) Error() string {
	return fmt.Sprintf("invalid radius is : %.2f, %s", e.radius, e.msg)
}

func circleArea(r float64) (float64, error) {
	if r < 0 {
		return 0, &areaErr{radius: r, msg: "please input valid radius"}
	} else {
		return math.Pi * r * r, nil
	}
}
