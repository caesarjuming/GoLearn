package main

import "fmt"

func main() {
	var b bool = true
	fmt.Println(b)
	const i int8 = 100
	fmt.Println(i)
	// 16 进制
	const i1 int16 = 0x111
	// 8 进制
	const i2 int16 = 011

	fmt.Println(i1, i2)

	// byte == uint8
	// rune == uint32

	// int uint 字节是变的

	const f1 float64 = 10
	const f2 float32 = 11.11
	const f3 float32 = .123

	fmt.Println(f1, f2, f3)

	const com1 complex64 = 1 + 2i
	const com2 complex128 = 1 - 2i

	fmt.Println(com1)
	fmt.Println(com2)

	// unicode
	const r rune = 'A'
	fmt.Println(r)

	const hz rune = '好'
	fmt.Println(hz)

	//ASCII
	const asc rune = '\xff'
	// 0-255 3个8位数
	const mr rune = '\111'
	// 4
	const ur rune = '\uffff'
	// 8
	const ulr rune = '\U00000000'
	fmt.Println(asc, mr, ur, ulr)

}
