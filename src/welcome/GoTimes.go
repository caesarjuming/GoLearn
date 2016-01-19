package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second)
	<-time1.C
	fmt.Println("过期")

	time2 := time.NewTimer(time.Second * 2)
	go func() {
		<-time2.C
		fmt.Println("time2过期")
	}()
	//提前取消
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("time2 stop")
	}

}
