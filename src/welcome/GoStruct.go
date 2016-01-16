package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{name: "aaa", age: 26})
	fmt.Println(person{"bbb", 55})
	//可以省略参数，为0值
	fmt.Println(person{name: "ccc"})

	//p是一个指针
	p := &person{name: "c", age: 99}
	p.age = 0
	fmt.Println(p)

}
