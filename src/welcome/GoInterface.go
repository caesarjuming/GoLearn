package main

import (
	"fmt"
	"strconv"
)

type Man interface {
	say()
	walk(lens int)
}
type One struct {
	name string
	age  int
}
type Two struct {
	habbit string
	play   string
}

//实现接口的方法， 实现接口其中所有的方法就是实现了这个接口
func (o One) say() {
	fmt.Println(o.name + strconv.Itoa(o.age))
}
func (t Two) say() {
	fmt.Println(t.habbit + t.play)
}
func (o One) walk(i int) {
	fmt.Println(i)
}
func (t Two) walk(i int) {
	fmt.Println(i + 100)
}

func main() {
	a := One{"jack", 28}
	b := Two{"football", "spring"}

	dd(a)
	dd(b)

}

func dd(m Man) {
	m.say()
	m.walk(1)
}
