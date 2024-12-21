package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func is running...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func is running...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}

func main() {
	// defer fmt.Println("main is over!!")

	// fmt.Println("main is running...")
	returnAndDefer()
}
