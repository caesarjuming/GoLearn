package main

import (
	"fmt"
	"time"
)

func main() {

	//默认的channel是没有缓冲的
	messages := make(chan string)
	go func() {
		messages <- "输入一个消息"
	}()

	rev := <-messages
	fmt.Println(rev)

	//带缓冲的channel,2个缓冲
	bufferMessages := make(chan string, 2)
	bufferMessages <- "哈哈"
	bufferMessages <- "呵呵"
	v1 := <-bufferMessages
	v2 := <-bufferMessages
	fmt.Println(v1, v2)

	//同步
	asyncMessages := make(chan string, 1)
	worker(asyncMessages)
	<-asyncMessages

}

func worker(s chan string) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	s <- "done"
}
