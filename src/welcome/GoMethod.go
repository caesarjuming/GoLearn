package main

import (
	"fmt"
)

type product struct {
	name  string
	price int
}

//结构体方法
func (p product) howMany() int {
	return p.price
}

//结构体指针方法
func (p *product) whatName() string {
	return p.name
}

func main() {
	p := product{name: "book", price: 50}
	fmt.Println(p.howMany())
	fmt.Println(p.whatName())
	pn := &p

	fmt.Println(pn.howMany())
	fmt.Println(pn.whatName())

}
