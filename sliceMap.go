package main

import "fmt"

func main() {
	s:=make([]map[string]string,0,3)
	m:=map[string]string{"a":"1","b":"2"}

	for i := 0; i < cap(s); i++ {
		s=append(s,m)
	}

	fmt.Println(s )
}
