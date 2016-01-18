package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops uint64 = 0
	for i := 0; i < 100; i++ {
		total := 0
		for {
			go func() {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}()
		}
	}

	for j := 0; j < 5; j++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("opsFinal", opsFinal)
	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()

}
