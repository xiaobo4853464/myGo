package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a:="helloworld"
	fmt.Println(strings.Contains(a,"h"))
	fmt.Println(strings.ContainsAny(a,"abch"))// 只要包含一个就true
	fmt.Println(strings.HasPrefix(a,"hello")) // 前缀
	fmt.Println(strings.HasSuffix(a,"world")) // 后缀
	fmt.Println(strings.Index(a,"h")) // 某个字符的index

	s:=[]string{"s","h","a","w"}
	words:=strings.Join(s,"") // 拼接
	fmt.Println(words)

	sp:=strings.Split(words,"") //分割
	fmt.Println(sp)

	a1:=strings.Replace(a,"o","0",-1) // 替换所有
	fmt.Println(a1)

	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))

	fmt.Println(a[5:]) //截取

	fmt.Println("-------------分割线------------")

	t,_:=strconv.ParseBool("true")
	fmt.Println(t)
	// str转 int
	aa,_:=strconv.ParseInt("111",10,64)
	fmt.Println(aa)
	fmt.Printf("%T\n",aa)

	// int 转str
	v:=strconv.FormatInt(aa,10)
	fmt.Println(v)
	fmt.Printf("%T\n",v)



}
