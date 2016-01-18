package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	//原子操作
	var ops uint64 = 0
	for i := 0; i < 3; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				fmt.Println(atomic.LoadUint64(&ops))
				runtime.Gosched()
			}

		}()
	}
	time.Sleep(time.Millisecond)
	finalOps := atomic.LoadUint64(&ops)
	fmt.Println("final:", finalOps)

}
