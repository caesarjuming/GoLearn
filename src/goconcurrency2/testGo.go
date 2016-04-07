package main

import "fmt"
import gc "goconcurrency"

// 只执行包初始化，不需要引用其中的函数变量
import _ "runtime/cgo"

func main() {
	gc.Hello()
	fmt.Println(gc.A)
	fmt.Println(gc.B)

	fmt.Println(gc.Person{Name: "aaa", Age: 18})
	// 注意：int 和int8不是同一类型
	// var a int = uint8(8)
	// cannot use uint8(8) (type uint8) as type int in assignment
	fmt.Println(a)

}
