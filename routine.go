package main

import (
	"fmt"
	"time"
)

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// go routine
func main() {
	go counter()
	go counter()
	counter()
}
