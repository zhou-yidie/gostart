package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	var a int = 10
	var b int = 20

	fmt.Println("a = ", a, " b = ", b)
	swap(&a, &b)
	fmt.Println("After swap:")
	fmt.Println("a = ", a, " b = ", b)
}
