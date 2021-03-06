package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 3)
	c2 := make(chan string, 3)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "c1 resource"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 2")
	}

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "c2 resource"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 3")
	}

}
