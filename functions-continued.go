package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("4 + 5 = ", add(4, 5))
}
