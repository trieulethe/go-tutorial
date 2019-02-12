package main

import "fmt"

// 0: even
// 1: odd
func checkOddOrEven(value int) int {
	return value & 1
}

func main() {
	res := 5 ^ 5
	fmt.Println("res:", res)
}
