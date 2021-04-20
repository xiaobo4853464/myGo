package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {

		a := rand.Intn(100)
		fmt.Println(a)

	}

}
