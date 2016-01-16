package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2)
	fmt.Println(s)

	a, b := doubleReturn()
	fmt.Println(a, b)

	fmt.Println(varFunction(1, 2, 3, 4, 5))
	//...把数组分解为多个参数，传递
	fmt.Println(varFunction([]int{1, 2, 3}...))

	//闭包，里面的值递增
	f := anonFunction()
	fmt.Println(f())
	fmt.Println(f())

	//递归
	fmt.Println(recurFunction(5))

	//指针
	p := 100
	//传递是指针，值被改变
	pointFunction(&p)
	fmt.Println(p)

}

//定义函数，多个参数可以用逗号隔开，返回值必须用return，不能省略
func sum(i, j int) int {
	return i + j
}

//多返回值
func doubleReturn() (int, int) {
	return 1, 2
}

//变参函数
func varFunction(value ...int) int {
	num := 0
	for _, v := range value {
		num += v
	}
	return num
}

//一个函数，返回值是一个返回值为int的匿名函数,可以形成闭包
func anonFunction() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//递归
func recurFunction(i int) int {
	if i > 1 {
		return i + recurFunction(i-1)
	} else {
		return 1
	}
}

//指针
func pointFunction(i *int) {
	*i = 987
}
