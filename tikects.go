package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var tickets = 10
var M sync.Mutex  //锁
var wg_ sync.WaitGroup //同步等待组
func main() {
	wg_.Add(4)
	go saleTickets("窗口1")
	go saleTickets("窗口2")
	go saleTickets("窗口3")
	go saleTickets("窗口4")
	wg_.Wait()
	//time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	defer wg_.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		M.Lock()
		if tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, " 卖出：", tickets)
			tickets--
		} else {
			M.Unlock()
			fmt.Println(name,": 没票了")
			break
		}
		M.Unlock()
	}
}
