package main

import (
	"fmt"
	"sync"
	"time"
)

var RW *sync.RWMutex
var wgg *sync.WaitGroup

func main() {
	RW = new(sync.RWMutex)
	wgg = new(sync.WaitGroup)
	wgg.Add(5)
	go write(1)
	go read(1)
	go write(2)
	go read(2)
	go write(3)

	wgg.Wait()
	//println()
}

func read(n int) {
	defer wgg.Done()
	fmt.Println(n, "开始读")

	RW.RLock()
	fmt.Println(n, "读。。。。")
	RW.RUnlock()
	fmt.Println(n, "读完成")
	time.Sleep(1 * time.Second)
}

func write(n int){
	defer wgg.Done()
	fmt.Println(n, "开始写")

	RW.Lock()
	fmt.Println(n, "写。。。。")
	time.Sleep(1 * time.Second)
	RW.Unlock()
	fmt.Println(n, "写完成")

}