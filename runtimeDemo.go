package main

import (
	"fmt"
	"runtime"
)

func main() {
	//设置go 程序执行的最大cpu数量 1-256
	//n := runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println(n)

	fmt.Println(runtime.GOROOT())
	// 获得操作平台
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())


	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine...")
		}
	}()

	for i := 0; i < 5; i++ {
		runtime.Gosched() //让出时间片，让其他 goroutine 先执行
		fmt.Println("main...")
	}

}
