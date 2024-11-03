package main

import (
	"fmt"
	"net/http"
)

func test() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("hello,go")
}
