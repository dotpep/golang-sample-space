package main

import (
	"errors"
	"fmt"
)

func sub(x int, y int) int {
	return x - y
}

func concat(s1, s2 string) string {
	return s1 + s2
}

func increment(x int) int {
	x++
	return x
}

func getNames() (string, string) {
	return "John", "Doe"
}

func getCoordsFirstSameVersion() (x, y int) {
	// x and y are initialized with zero values
	return // automatically returns x and y
}

func getCoordsSecondSameVersion() (int, int) {
	var x, y int
	return x, y
}

func getCoordsThirdSameVersion() (x, y int) {
	return x, y
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	divideResult := dividend / divisor

	return divideResult, nil
}

type insuranceStatus interface {
	hasInsurance() bool
	isTotaled() bool
	isDenied() bool
	isBigDent() bool
}

func getInsiranceAmount(status insuranceStatus) int {
	// avoid using nested conditional statements instead of use linear like aproach
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDenied() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}

func main() {
	//result := sub(5, 2)
	result := concat("Hello ", "World!")
	fmt.Printf("result: %v \n", result)

	fmt.Println("---")

	// In Golang integer type variable and using function to increment it values are Primitive/Value type not Reference
	// and we must reasign incremented value to x variable
	x := 5
	x = increment(x)
	fmt.Printf("incremented: %v \n", x)

	fmt.Println("---")

	firstName, secondName := getNames()
	fmt.Printf("firstName: %v, secondName: %v \n", firstName, secondName)

	_, lastName := getNames()
	fmt.Printf("lastName: %v \n", lastName)

	fmt.Println("---")

	// Early returns
	divideResult, err := divide(10, 5)
	fmt.Printf("first test result: %v, error: %v \n", divideResult, err)

	divideResult, err = divide(10, 0)
	fmt.Printf("second test result: %v, error: %v \n", divideResult, err)
}
