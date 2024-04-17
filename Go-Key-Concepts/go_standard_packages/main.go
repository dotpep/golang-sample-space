package main

import (
	"fmt"
	"math"
)

func main() {
	//sum := addr(1, 2)
	//fmt.Println(sum)
	conditionsExample()
}

func addr(a, b int) int {
	return a + b
}

func argFunc(myFunc func(int, int) int) int {
	return myFunc(1, 2)
}

func conditionsExample() {
	fmt.Println("true && false statement is: ", true && false)
	fmt.Println("true || false statement is: ", true || false)
	fmt.Println("!true statement is: ", !true)

	if true && false {
		fmt.Println("or && logical operator returns true")
	}
	if true || false {
		fmt.Println("and || logical operator returns true")
	}
	if !false {
		fmt.Println("is not false, is true")
	}

	fmt.Println("---")

	num := 8
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	if num%4 == 0 {
		fmt.Printf("%d is divisible by 4\n", num)
	}

	fmt.Println("---")

	if myNum := 100; myNum < 0 {
		fmt.Printf("%d is negative\n", myNum)
	} else if myNum < 10 {
		fmt.Printf("%d has 1 digit\n", myNum)
	} else {
		fmt.Printf("%d has multiple digit\n", myNum)
	}

	fmt.Println("---")

	switch myNum := 2; myNum {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func variablesExample() {
	//num := 5
	//var num2 = 15
	var num3 uint = 1
	var num3float float64 = -2.70
	var num4 uint = uint(math.Abs(num3float))
	//fmt.Println(num + num2 + num3)
	fmt.Println(num3 + num4)

	newNum := 3.05
	fmt.Printf("type of variable is: %T\nand is value is: %v\n", newNum, newNum)

	// byte is just alias for uint8
	var num5 byte = 100
	var num6 uint8 = 100
	fmt.Println(num5, num6)

	// rune is alias for int32
	var num7 rune = 10
	fmt.Println(num7)

	var num10 complex128 = 11
	fmt.Println(num10)

	var isBool bool = true
	fmt.Println(isBool)

	var text string = "Hello, world!"
	fmt.Println(text)

	const firstName = "John"
	const lastName = "Doe"
	const fullName = firstName + " " + lastName
	fmt.Println(fullName)

	const secondsInMinute, minutesInHour = 60, 60
	const secondsInHour = minutesInHour * secondsInMinute
	fmt.Println(secondsInMinute, minutesInHour, secondsInHour)

	// Formating
	// %v - Interpolate default representation
	fmt.Printf("I am %v years old\n", 10)
	fmt.Printf("I am %v years old\n", "way too many")

	// pretty custom struct type formating
	// is anonymous struct
	emp := struct {
		name   string
		age    int
		salary int
	}{name: "Sam", age: 31, salary: 2000}

	fmt.Printf("%v", emp)  // result (only values): {Sam 31 2000}
	fmt.Printf("%+v", emp) // result (field and value): {name:Sam age:31 salary:2000}
	fmt.Printf("%#v", emp) // result (struct name and both): main.employee{name:"Sam", age:31, salary:2000}

	// %s - Interpolate string
	fmt.Printf("I am %s years old\n", "way too many")

	// %d - Interpolate integer in decimal form
	fmt.Printf("I am %d years old\n", 10)

	// %f - Interpolate decimal
	fmt.Printf("I am %f years old\n", 10.523)
	fmt.Printf("I am %.2f years old\n", 10.523)
	fmt.Printf("I am %.f years old\n", 10.523)

	// %T - Display type of value
	fmt.Printf("I am %T years old\n", 10)

	// Print memory address of variable
	a := 2
	fmt.Println(&a)
	// %p - Print address of variable (using formating)
	fmt.Printf("The address of a is: %p\n", &a)

	// Sprintf for creating formatted strings without printing them directly (output and return a string)
	const name = "Adam"
	const openRate = 30.5
	msg := fmt.Sprintf("Name: %s, Open rate: %.2f", name, openRate)
	println(msg)

	// Scientific notation
	const d = 3e20
	formatted := fmt.Sprintf("%.0f", d)
	fmt.Printf("scientific notation: %v, formatted version: %v, count of digit: %v\n", d, formatted, len(formatted))

}
