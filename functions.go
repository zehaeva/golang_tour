package main

import (
	"fmt"
)

func add(int x, int y) int {
	return x + y
}

func main() {
	fmt.Println("4 + 5 = ", add(4, 5))
}
