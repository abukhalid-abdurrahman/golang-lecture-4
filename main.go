package main

import (
	"fmt"
)

func div (a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func multi(a int, b int) int {
	res := a * b
	return res
}

func main() {
	divRes := div(4, 2)
	multiRes := multi(5, 2)
	fmt.Println(multiRes)
	fmt.Println(divRes)
}