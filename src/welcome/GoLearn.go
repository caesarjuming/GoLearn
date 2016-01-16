package main

import (
	"fmt"
	"time"
)

//https://gobyexample.com/slices

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())

	//分别赋值
	var a, b int = 1, 2
	fmt.Println(a, b)
	//不需要明确类型，自动推导
	var c = true
	fmt.Println(c)
	//默认的值为0值,如果没有初始赋值，必须指定类型
	var d int
	fmt.Println(d)
	//变量声明和类型声明初始化简写
	e := 3
	fmt.Println(e)

	//定义常量，自动推断类型
	const f = 111
	fmt.Println(f)
	//声明类型
	const g int = 222
	fmt.Println(g)
	//强制转换
	fmt.Println(int64(g))

	//常见for循环,注意：没括号
	for i := 0; i < 10; i++ {

	}
	for true {
		fmt.Println("单一条件")
		break
	}
	for {
		//一直执行
		break
	}

	//if可以初始化变量
	if num := 7; num < 10 {
		fmt.Println(num, "num")
	} else if num > 10 {

	} else {

	}

	//switch
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}

	//用逗号来分割多个条件
	ii := 2
	switch ii {
	case 1, 2, 3:
		fmt.Println("1,2,3")
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("default")
	}

	//在外面定义条件
	var bb = 1 + 2

	switch {
	case bb > 1:
		fmt.Println("大于1")
	default:
		fmt.Println("default")
	}

	//定义数组,初始值都是0
	var arr [10]int
	fmt.Println(arr[1])
	arr[2] = 333
	fmt.Println(arr)
	fmt.Println(arr[2])

	arr2 := [10]int{1, 2, 3, 4, 5}
	fmt.Println(arr2[5])
	fmt.Println(arr2[4])
	fmt.Println(arr2, len(arr2))

	doubleArr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(doubleArr)
	fmt.Println(doubleArr[0][0])

	//数组切片,声明变量，不使用会报错
	sliceArr := make([]int, 3)
	fmt.Println(sliceArr)

	sliceArr[0] = 100
	fmt.Println(len(sliceArr))
	sliceArr = append(sliceArr, 200, 300, 400)
	fmt.Println(sliceArr)

	//切片是指针传递
	ssslice := sliceArr
	ssslice[0] = 9999
	fmt.Println(sliceArr)

	//切片操作和python类似
	s1 := sliceArr[0:3]
	fmt.Println(s1)
	s2 := sliceArr[:]
	fmt.Println(s2)
	//构造一个多维不对齐的切片数组
	supers := make([][]int, 3)
	for i := 0; i < 3; i++ {
		supers[i] = make([]int, i+1)
	}
	fmt.Println(supers)

	//Maps 创建一个String，int的Map
	maps := make(map[string]int, 5)
	maps["aaa"] = 2
	maps["bbb"] = 3

	//删除一个key,value
	delete(maps, "aaa")
	fmt.Println(maps)

	value, ifGetSuccess := maps["bbb"]
	fmt.Println(ifGetSuccess, value)

	//另一种方式构建
	newMaps := map[string]int{"a": 1, "b": 2}
	fmt.Println(newMaps)

	all := 0
	numbers := []int{1, 2, 3, 4, 5}
	//对于第一个变量，如果不需要也可以这样 for _,num :=range numbers
	for i, num := range numbers {
		all += num
		fmt.Println(i)
	}
	fmt.Println(all)

	// for map

	for k, v := range map[string]string{"a": "!", "b": "@", "c": "#"} {
		fmt.Println(k, v)
	}
}
