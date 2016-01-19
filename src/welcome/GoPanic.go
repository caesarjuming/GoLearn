package main

import (
	"os"
)

func main() {
	panic("wrong!")
	_, err := os.Create("C:\\abc.txt")
	if err != nil {
		panic(err)
	}

}
