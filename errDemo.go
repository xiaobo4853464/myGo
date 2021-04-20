package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//age:=-3
	//err:=assertAge(age)
	//if err!=nil{
	//	fmt.Println(err)
	//}

	_,err:=os.Open("xiaobo.text")
	if ins,ok:=err.(*os.PathError);ok{
		fmt.Println(ins.Op)
		fmt.Println(ins.Path)
		fmt.Println(ins.Err)
	}

	//hostName:="www.google.com"
	hostName:="www.google1112234.com"
	s,e:=net.LookupHost(hostName)
	if e!=nil{
		if ins,ok:=e.(*net.DNSError);ok{
			fmt.Println(ins.Name)
			fmt.Println(ins.Server)
		}
		//fmt.Println(e)
	}
	fmt.Println(s)

}

func assertAge(age int) error {
	if age < 0 {
		//e:=errors.New("年龄不合法")
		e := fmt.Errorf("%d,年龄不合法", age)
		return e
		//return
	} else {
		fmt.Println("年龄合法")
		return nil
	}
}
