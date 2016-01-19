package main

import (
	"fmt"
)

func onlyIn(c chan<- string, msg string) {
	c <- msg
}
func inOut(i <-chan string, o chan<- string) {
	temp := <-i
	o <- temp
}

func main() {
	onlyInP := make(chan string, 3)
	inOutP := make(chan string, 3)

	onlyIn(onlyInP, "传入")
	fmt.Println(<-onlyInP)

	onlyIn(onlyInP, "传入2")
	inOut(onlyInP, inOutP)
	fmt.Println(<-inOutP)

}
