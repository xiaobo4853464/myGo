package main

import "fmt"

func main() {
	m:=map[string]int {"a":1,"b":2}
	var m1 map[string]int // 只是初始化
	var m2=make(map[string]int)

	fmt.Println(m==nil)
	fmt.Println(m1==nil)
	fmt.Println(m2==nil)
	//fmt.Printf("%v\n", m)
	for i, i2 := range m {
		println(i,i2)
	}

	// 判断k是否存在
	k,ok:=m["a"]
	if ok{
		println(k)
	}

	delete(m,"a")
	fmt.Printf("%v", m)




}
