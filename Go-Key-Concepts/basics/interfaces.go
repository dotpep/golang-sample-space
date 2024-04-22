package main

import (
	"fmt"
	"math"
	"time"
)

// Interface is contract in implementation
// Interface name convention must ends with 'er'
// or starts with 'I' and UpperCase
type Shaper interface {
	area()
	perimeter()
}

type rectangle struct {
	width, height float64
}

func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

type circle struct {
	radius float64
}

func (crcl circle) area() float64 {
	return math.Pi * crcl.radius * crcl.radius
}

func (crcl circle) perimeter() float64 {
	return 2 * math.Pi * crcl.radius
}

// Interface Example 1
type Messager interface {
	getMessage() string
}

func sendMessage(msg Messager) {
	fmt.Println(msg.getMessage())
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bdmsg birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bdmsg.recipientName, bdmsg.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sndrep sendingReport) getMessage() string {
	return fmt.Sprintf("Your '%s' report is ready. You've sent %v", sndrep.reportName, sndrep.numberOfSends)
}

func testInterface1(msg Messager) {
	sendMessage(msg)
	fmt.Println("==================")
}

// Interface Example 2
type Employeer interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (cont contractor) getName() string {
	return cont.name
}

func (cont contractor) getSalary() int {
	return cont.hoursPerYear * cont.hourlyPay
}

type fullTime struct {
	name   string
	salary int
}

func (fullt fullTime) getName() string {
	return fullt.name
}

func (fullt fullTime) getSalary() int {
	return fullt.salary
}

func testInterface2(emp Employeer) {
	fmt.Printf("Name: %s, \nSalary per year: %v\n", emp.getName(), emp.getSalary())
	fmt.Println("==================")
}

// Multiple Interfaces Example 3
type Expenser interface {
	cost() float64
}

type Printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

func (em email) cost() float64 {
	msgSize := float64(len(em.body))

	if !em.isSubscribed {
		return msgSize * 0.05
	}
	return msgSize * 0.01
}

func (em email) print() {
	fmt.Println(em.body)
}

func print(prnt Printer) {
	prnt.print()
}

func testMultiInterface3(exp Expenser, prnt Printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", exp.cost())
	prnt.print()
	fmt.Println("==================")
}

// Better Interface Definition for:
//
//	type Copier interface {
//		Copy(string, string) int
//	}
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// Type Assertions
//type shape interface { area() }
//type circle struct { radius float64 }

// 'c' is a new circle cast from 's'
// which is an instance of a shape.
// 'ok' is a bool that is true if s was a circle
// or false if s isn't a circle
//c, ok := s.(circle)

// Type Assertions Example 4
func getExpenseReport(expense Expenser) (string, float64) {
	mailMsg, ok := expense.(email)
	if ok {
		return mailMsg.toAddress, mailMsg.cost()
	}
	sMsg, ok := expense.(sms)
	if ok {
		return sMsg.toPhoneNumber, sMsg.cost()
	}
	return "", 0.0
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (s sms) cost() float64 {
	msgSize := float64(len(s.body))

	if !s.isSubscribed {
		return msgSize * 0.1
	}
	return msgSize * 0.03
}

func (i invalid) cost() float64 {
	return 0.0
}

// Type Switches
func printNumericValue(num interface{}) {
	switch val := num.(type) {
	case int:
		fmt.Printf("%T\n", val)
	case string:
		fmt.Printf("%T\n", val)
	default:
		fmt.Printf("%T\n", val)
	}
}

// printNumericValue(1) // prints 'int'
// printNumericValue("1") // prints 'string'
// printNumericValue(struct{}{}) // prints 'struct {}'

// Type Switches Example 4
func getSwitchedExpenseReport(expense Expenser) (string, float64) {
	switch valT := expense.(type) {
	case email:
		return valT.toAddress, valT.cost()
	case sms:
		return valT.toPhoneNumber, valT.cost()
	default:
		return "", 0.0
	}
}

func testTypeAssertionAndSwitcher4(exps Expenser) {
	// replace getExpenseReport(exps) to getSwitchedExpenseReport(exps)
	// for testing with Switching instead Assertion
	address, cost := getExpenseReport(exps)
	switch exps.(type) {
	case email:
		fmt.Printf("Report: The EMAIL going to %s will cost: $%.2f ...\n", address, cost)
		fmt.Println("==================")
	case sms:
		fmt.Printf("Report: The SMS going to %s will cost: $%.2f ...\n", address, cost)
		fmt.Println("==================")
	default:
		fmt.Printf("Report: Invalid expense\n")
		fmt.Println("==================")
	}
}

// Write Interfaces cleaner
// 1. Keep Interfaces Small as you can
// 2. Interfaces Should Have No Knowledge of Satisfying Types
// 3. Interfaces Are Not Classes
// 10. Apply Naming Convention of 'er' naming ends and UpperCase or 'I' prefix on starts

func main() {
	// Interfaces Example 1
	testInterface1(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	testInterface1(birthdayMessage{
		birthdayTime:  time.Date(1999, 03, 21, 0, 0, 0, 0, time.UTC),
		recipientName: "Ashley",
	})
	testInterface1(sendingReport{
		reportName:    "Second Report",
		numberOfSends: 50,
	})
	testInterface1(birthdayMessage{
		birthdayTime:  time.Date(2000, 04, 20, 0, 0, 0, 0, time.UTC),
		recipientName: "John",
	})

	fmt.Println("---")

	// Interfaces Example 2
	testInterface2(fullTime{
		name:   "Alex",
		salary: 100000,
	})
	testInterface2(contractor{
		name:         "John",
		hoursPerYear: 650,
		hourlyPay:    120,
	})
	testInterface2(contractor{
		name:         "Bob",
		hoursPerYear: 480,
		hourlyPay:    180,
	})
	testInterface2(fullTime{
		name:   "Jack",
		salary: 80000,
	})

	fmt.Println("---")

	// Multiple Interfaces Example 3
	mail := email{
		isSubscribed: true,
		body:         "Hello there",
	}
	testMultiInterface3(mail, mail)
	mail = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	testMultiInterface3(mail, mail)

	fmt.Println("---")

	// Type Assertions Example 4
	testTypeAssertionAndSwitcher4(email{
		isSubscribed: true,
		body:         "Hello there",
		toAddress:    "john@does.com",
	})
	testTypeAssertionAndSwitcher4(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+15554441010",
	})
	testTypeAssertionAndSwitcher4(invalid{})
}
