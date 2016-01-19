package main

import (
	"fmt"
)

func testThread(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, i)
	}
}

func main() {
	testThread("直接执行")
	go testThread("线程执行")

	//执行一个匿名函数的线程
	go func(s string) {
		fmt.Println("我在执行匿名函数：" + s)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
