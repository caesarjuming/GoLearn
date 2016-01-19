package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 3)
	c2 := make(chan string, 3)

	go func() {
		c1 <- "One"
		time.Sleep(time.Second * 1)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}

}
