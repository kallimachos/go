package main

import (
	"fmt"
	"net/http"
)

func get(c chan<- *http.Response) {
	resp, _ := http.Get("http://example.com/")
	c <- resp
}

func main() {
	c := make(chan *http.Response)
	for i := 0; i < 10; i++ {
		go get(c)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
