package main

import "fmt"

func main() {

	var s string ="xiaobo"
	fmt.Printf("%T,%s\n",s,s)

	s1:=`helloworld`
	fmt.Printf("%T,%s\n",s1,s1)

	// 区别
	v1:='a'
	v2:="a"
	fmt.Printf("%T,%d\n",v1,v1)
	fmt.Printf("%T,%s\n",v2,v2)

}
