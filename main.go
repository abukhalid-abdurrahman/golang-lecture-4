package main

import (
	"fmt"
)

func multi(a int, b int) int {
	res := a * b
	return res
}

func main() {
	multiRes := multi(5, 2)
	fmt.Println(multiRes)
}