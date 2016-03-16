package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "I can't type chinese into vim"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	// strange this was yesterday!

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
