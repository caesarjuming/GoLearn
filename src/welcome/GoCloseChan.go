package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	//这里的done和锁类似
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("rec", j)
			} else {
				fmt.Println("all rev")
				done <- true
				return
			}
		}

	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send", i)
	}
	close(jobs)
	fmt.Println("send all")
	<-done

	//迭代chan
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	//如果不关闭的话，那么迭代的时候就会阻塞在获取第三个
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}

}
