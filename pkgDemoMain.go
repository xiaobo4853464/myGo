package main

import (
	"MyGoProject/pkgDemo"
	"fmt"
)

func main() {
	pkgDemo.GetDbConn()
	fmt.Println("演示了如何使用pkg")

}
