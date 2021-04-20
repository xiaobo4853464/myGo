package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, e := rpc.DialHTTP("tcp", "localhost:8089")
	if e != nil {
		panic(e)
	}

	var req float64 = 5
	var resp *float64

	methodName := "MathUtil.CalArea"
	//同步调用
	//err := client.Call("MathUtil.CalArea", req, &resp)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(*resp)

	//异步调用
	syncCall := client.Go(methodName, req, &resp, nil)
	respDone := <-syncCall.Done
	fmt.Println(respDone)
	fmt.Println(*resp)

	//for i := 0; i < 100; i++ {
	//	go tt(client,float64(i))
	//}
	//time.Sleep(3*time.Second)

}

func tt(client *rpc.Client,r float64){
	var resp *float64
	syncCall := client.Go("MathUtil.CalArea", r, &resp, nil)
	respDone := <-syncCall.Done
	fmt.Println(respDone)
	fmt.Println(*resp)
}