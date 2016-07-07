package main

import (
	"fmt"
)

type binFunc func(...int) int

func add(args ...int) int {
	if len(args) > 0 {
		return args[0] + add(args[1:]...)
	}
	return 0
}

func mul(args ...int) int {
	if len(args) > 0 {
		return args[0] * mul(args[1:]...)
	}
	return 1
}

func (f binFunc) Error() string {
	return "binFunc error: using without method"
}

func (f binFunc) Calc(args ...int) int {
	return f(args...)
}

func main() {
	fmt.Println(binFunc(add).Calc(1, 2, 3, 4, 5))
	fmt.Println(binFunc(mul).Calc(1, 2, 3, 4, 5))
	fmt.Println(binFunc(add))
}
