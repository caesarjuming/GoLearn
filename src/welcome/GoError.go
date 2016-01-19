package main

import (
	"errors"
	"fmt"
)

func testError(i int) (int, error) {
	if i > 10 {
		fmt.Println("大于10")
		return i, errors.New("不支持大于10")
	} else {
		fmt.Println("小于等于10")
		return i, nil
	}
}

//定制errors
type myError struct {
	count  int
	reason string
}

func (m myError) Error() string {
	fmt.Println(m.reason)
	return m.reason
}

func f2(i int) (int, error) {
	if i > 5 {
		return i, myError{8, "不能大于5"}
	}
	return i, nil
}

func main() {
	testError(100)
	fmt.Println(f2(8))
}
