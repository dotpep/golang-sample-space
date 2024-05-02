package main

import (
	"fmt"

	"github.com/dotpep/mycalculator"
	"github.com/dotpep/mystrings"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(
		mystrings.Reverse("hello"),
	)
	fmt.Println(
		mycalculator.Add(2, 3),
	)
	fmt.Println(
		mycalculator.Multiply(5, 5),
	)
}
