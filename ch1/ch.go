package main

import "fmt"

func ch() {
	ch := make(chan string)

	go func() {
		ch <- "Hello, channel!"
	}()

	fmt.Println(<-ch)
}
