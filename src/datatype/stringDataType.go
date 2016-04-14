package main

import "fmt"

func main() {

	//raw string
	var s string = `"hello",\\aa`

	var s1 string = "hello,\t \\"
	fmt.Println(s, s1)
	// string 都是不可变的

	// 数组的长度也是数组的属性，两个长度不同的同类型的数组不是同一类型
	a2 := [2]string{}
	a1 := [3]string{}

	//一个匿名结构体的数组
	a3 := [5]struct{ name, address string }{}

	fmt.Println(a1, a2, a3)

}
