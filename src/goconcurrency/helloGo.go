package goconcurrency

import "fmt"

// 包声明语句不算声明，因为不在任何作用域内

// 包初始化函数
func init() {
	helloW := "hello"
	fmt.Println(helloW)
}

// 因为是大写，所以可以导出
var A string = "AAAAAAAA"

const B string = "BBBBBBBBBB"

func main() {
	// GO 由Unicode编码组成，go语言源码必须是UTF-8
	fmt.Print("aaa")
	// fmt.Println(helloW)  不能引用
	// 只有首字母大写的代码块才是可包导出，在包外面可见

}

// 此函数是可以导出的
func Hello() {
	fmt.Println("hello,Out")
}

type Person struct {
	Name string
	Age  uint8
}
