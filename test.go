package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("hello,go")
}
