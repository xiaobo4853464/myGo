package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	m := new(MathUtil)
	a := new(AddParam)

	// 注册服务
	registerErr := rpc.Register(m)
	addErr := rpc.Register(a)
	if registerErr != nil {
		fmt.Println(registerErr.Error())
	} else if addErr != nil {
		fmt.Println(addErr.Error())
	}
	//把服务注册在http协议上，可以用http 进行数据传递
	rpc.HandleHTTP()
	listen, listenErr := net.Listen("tcp", ":8089")
	//listen, listenErr := net.Listen("tcp", ":8090")
	if listenErr != nil {
		fmt.Println(listenErr.Error())
	}
	http.Serve(listen, nil)

}

type MathUtil struct {
}

type AddParam struct {
	a int
	b int
}

func (m *MathUtil) CalArea(req float64, resp *float64) error {
	//var resp float64
	*resp = math.Pi * math.Pow(req, 2)
	return nil
}

func (a *AddParam) Add(req AddParam, resp *int) error {
	*resp = req.a + req.b
	return nil

}
