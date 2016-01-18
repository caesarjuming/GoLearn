package main

import (
	"fmt"
	"time"
)

func main() {

	//每隔一段时间执行一次
	myticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for i := range myticker.C {
			fmt.Println("@", i)
		}
	}()

	time.Sleep(time.Second * 5)
	myticker.Stop()
	fmt.Println("tickers stop")

}
