package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var inputAmount int

	fmt.Scanln(&inputAmount)

	var inputs []string

	for i := 0; i < inputAmount; i++ {
		var input string
		fmt.Scanln(&input)
		inputs = append(inputs, input)
	}

	for i := 0; i < inputAmount; i++ {
		fmt.Println(inputs[i])
	}
}
