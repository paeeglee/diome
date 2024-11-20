package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func(c chan string) {
		for {
			c <- "ping"
		}
	}(ch)
	go func(c chan string) {
		for {
			c <- "pong"
		}
	}(ch)

	go func(c chan string) {
		for {
			msg := <-c
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}(ch)

	var input string
	fmt.Scanln(&input)
}
