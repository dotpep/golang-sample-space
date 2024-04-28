package main

import (
	"fmt"
	"strings"
)

// Pointer
// if we have a varaible 'x:=5', 'y:=x' int
// our 'y' creates new address:value in RAM (as primitive data type)
// primitive type means that we cannot change value of original variable
// with assign it to new value in new variable that stores that value
// in case of using function we must reassign value of 'x'
// if we use it in func param and changed it
// otherwise it value cannot be changed

// poiters is fix this problem and make our 'x:=5'
// with defining 'z:=&x', ('var pointer *int')
// and reassign it using `*z=6` it will change value of 'x'
// it can be useful if we not need return any data in our function
// but we must change value of 'x' dynamically
// like we do with referenced data type (slices, maps) with append() build-in function

// x := 5
// y := x
// z := &x
// *z = 6

// Claster of Variable Name:Address:Value
// name | addr | value
// x    | 169  | 5 changed to 6
// y    | 170  | 5 (it copy value of 'x')
// z    | 171  | 169 (is address of 'x' in memory) mutable, connected to 'x' variable

// RAM
// addr | value
// 169  | 5 changed to 6 (because of 'z' pointer is referenced to 'x')
// 170  | 5 (is not be changed, 'y' is copy of previous 'x' value)
// 171  | 169 (is address of 'x' in memory) mutable, connected to 'x' variable
// ...  | ...

// in python we have mutable and unmutable data types
// in c# we have primitive and reference data types

// Pointer Example 1
func removeProfanity(message *string) {
	// nil Pointers
	// avoid is pointer setted as nil value (None)
	if message == nil {
		return
		//return errors.New("invalid input")
	}

	msgVal := *message
	msgVal = strings.ReplaceAll(msgVal, "dang", "**ng")
	msgVal = strings.ReplaceAll(msgVal, "shoot", "***ot")
	msgVal = strings.ReplaceAll(msgVal, "heck", "**ck")

	*message = msgVal
}

func testPointer1(messages []string) {
	for _, msg := range messages {
		removeProfanity(&msg)
		fmt.Println("changed: ", msg)
	}
}

// Pointer Receiver 1
type circle struct {
	x, y, radius int
}

func (c *circle) grow() {
	c.radius *= 2
}

// Pointer Receiver 2
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

// Pointer Example 2
// receiver

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func testPointer2(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage("this is my second draft")
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func main() {
	// Pointer

	// declare nil or 0, '' value pointer type of string
	var myStrPointer *string

	myString := "hello"
	myStrPointer = &myString

	// value before changes
	fmt.Println("myString: ", myString)
	fmt.Println("myStrPointer: ", myStrPointer)

	// change my hello to holla
	*myStrPointer = "holla"

	// value after changes
	fmt.Println("myString after changes of pointer: ", myString)
	fmt.Println("myStrPointer after changes of pointer: ", myStrPointer)

	// look up to de-referenced pointer value and address
	fmt.Println("*myStrPointer de-referenced value after changes: ", *myStrPointer)
	fmt.Println("adress of &*myStrPointer de-referenced value: ", &*myStrPointer)

	// address of value in RAM
	fmt.Printf("address of &myString: %p\n", &myString)
	fmt.Printf("address of &myStrPointer: %p\n", &myStrPointer)

	fmt.Println("---")

	// Pointer Example 1
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	testPointer1(messages1)
	testPointer1(messages2)

	fmt.Println("---")

	// Pointer Receiver 1
	circle1 := circle{
		x:      1,
		y:      2,
		radius: 4,
	}
	// notice 'c' is not a pointer in calling function
	// but method still gains access to a pointer to c
	fmt.Println("radius of circle1 before: ", circle1.radius)
	circle1.grow()
	fmt.Println("radius of circle1 after growth(): ", circle1.radius) // prints 8

	fmt.Println("---")

	// Pointer Receiver 2
	car1 := car{
		color: "white",
	}

	fmt.Println("color of car1 befor: ", car1)
	car1.setColor("blue")
	fmt.Println("color of car1 after: ", car1)

	fmt.Println("---")

	// Non-Pointer Receiver
	// cannot change dynamically
	// we must re-assign value of 'car1.setColor("blue")'

	// Pointer Example 2
	// receiver
	testPointer2(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

}
