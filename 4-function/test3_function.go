package main

import (
	"fmt"
)

func foo1(a string, b int) int {
	fmt.Println("a = ", a, " b = ", b)
	c := 100
	return c
}

func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a, " b = ", b)
	c, d := 100, 200
	return c, d
}

func main() {
	var a = "hello"
	var b = 100
	c := foo1(a, b)
	fmt.Println("c = ", c)
	ret1, ret2 := foo2(a, b)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
