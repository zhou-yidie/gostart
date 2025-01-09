package wg

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			//wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("main goroutine done!")
}
