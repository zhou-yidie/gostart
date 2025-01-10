package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main goroutine done!")
	time.Sleep(time.Second * 10)
}
