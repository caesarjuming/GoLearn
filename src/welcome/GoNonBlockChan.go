package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	select {
	case msg := <-c1:
		fmt.Println(msg)
	default:
		fmt.Println("no parm")
	}

	msg := "aaa"
	select {
	case c1 <- msg:
		fmt.Println("send a message", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msgRe := <-c2:
		fmt.Println("reci" + msgRe)
	case msgRe2 := <-c1:
		fmt.Println("reci2" + msgRe2)
	default:
		fmt.Println("nothing")
	}

}
