package main

import (
	"MyGoProject/protoPkg"
	"fmt"
	p "github.com/golang/protobuf/proto"
)

func main() {
	h := &protoPkg.Human{

		Name: "xiaobo",
		Age:  99,
	}
	fmt.Println(h)

	hEncoding, err := p.Marshal(h)
	if err != nil {
		panic(err)
	}
	fmt.Printf("after encoding:%v\n", hEncoding)

	hh := &protoPkg.Human{}
	err_ := p.Unmarshal(hEncoding, hh)
	if err_ != nil {
		panic(err_)
	}
	fmt.Println("after decoding:", hh)
	fmt.Println("name:", hh.GetName())
	fmt.Println("age:", hh.GetAge())

}
