package main

import "fmt"

func main() {
	num := 5
	switch num {
	case 1:
		fmt.Println("1st")
	case 2:
		fmt.Println("2nd")
	case 3:
		fmt.Println("3rd")
	case 4:
		fmt.Println("4th")
	default:
		fmt.Println("not found")
	}
	//fmt.Println("main...quit...")

	score:=100
	switch {
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	case score >= 60 && score < 80:
		fmt.Println("良好")
	case score >= 80 && score <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("输入有误")
	}

	fmt.Println("------------")

	switch m:=1;m {
	case 1:
		fmt.Println("1")
		fallthrough  // 穿透执行下一个
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("0")
	}
}
