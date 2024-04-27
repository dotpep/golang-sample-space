package main

import (
	"errors"
	"fmt"
)

// Advanced Function

// first class functions that can be treated as values
// funcs as data

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

// aggregate applies given math functuin to first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// Advanced Function Example 1
// First Class and Higher Order Functions
func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := make([]string, 0, len(messages))
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreating(message string) string {
	return "Hello! " + message
}

func testAdvanceFunc1(messages []string, formatter func(string) string) {
	defer fmt.Println("==================")
	formattedMessages := getFormattedMessages(messages, formatter)
	if len(formattedMessages) != len(messages) {
		fmt.Println("The number if messages returned is incorrect.")
		return
	}
	for i, message := range messages {
		formatted := formattedMessages[i]
		fmt.Printf(" * %s -> %s\n+", message, formatted)
	}
}

// Currying
// is used in middleware in backend side of web-dev
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

// Advanced Function Example 2

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(s1, s2 string) {
		fmt.Println(formatter(s1, s2))
	}
}

func testAdvanceFunc2(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("==================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimitFormatter(first, second string) string {
	return first + ": " + second
}

func commaDelimitFormatter(first, second string) string {
	return first + ", " + second
}

// Advanced Function Example 3
// Defer keyword
const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name    string
	number  int
	isAdmin bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	// this func should always delete user from user's map
	// and this 'defer' keyword
	// is will be executed 1 time
	// when this function complete own behavior, last time
	// in any returns case
	defer delete(users, name)

	user, ok := users[name]

	if !ok {
		return logNotFound
	}
	if user.isAdmin {
		return logAdmin
	}

	return logDeleted
}

func testAdvanceFunc3(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("==================")

	log := logAndDelete(users, name)
	fmt.Println("Log: ", log)
}

// Closures
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

// Advanced Function Example 4
type emailBill struct {
	costInPennies int
}

func adder() func(int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func testAdvanceFunc4(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

// Anonymous Function

// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, num := range nums {
		results = append(results, f(num))
	}
	return results
}

// Advanced Function Example 5

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(msg string) int {
			return len(msg)
		}, message)
	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf("Message: '%s' Cost: %v cents", message, cost)
	fmt.Println()
}

func testAdvanceFunc5(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func main() {
	// Advanced Function
	fmt.Println(aggregate(2, 3, 4, add))      // prints 9
	fmt.Println(aggregate(2, 3, 4, multiply)) // prints 24

	fmt.Println("---")

	// Advanced Function Example 1
	// First Class and Higher Order Functions

	// First Class funcs
	// func() int
	// func(string) int

	// Higher Order funcs
	// is function that takes a function as an argument or returns func as return value
	// func aggregate(a, b, c int, 'arithmetic func(int, int) int') int { }

	testAdvanceFunc1([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addSignature)
	testAdvanceFunc1([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addGreating)

	fmt.Println("---")

	// Currying
	// function carrying is practice of writing a func
	// that takes function (or functions) as input,
	// and returns a new function
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5)) // prints 25
	fmt.Println(doubleFunc(5)) // prints 10

	fmt.Println("---")

	// Advanced Function Example 2
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	testAdvanceFunc2("Error on database server", dbErrors, colonDelimitFormatter)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	testAdvanceFunc2("Error on mail server", mailErrors, commaDelimitFormatter)

	fmt.Println("---")

	// Advanced Function Example 3
	// Defer keyword

	usersMap := map[string]user{
		"john": {
			name:    "john",
			number:  12224443222,
			isAdmin: true,
		},
		"elon": {
			name:    "elon",
			number:  19875556452,
			isAdmin: true,
		},
		"breanna": {
			name:    "breanna",
			number:  98575554231,
			isAdmin: false,
		},
		"kade": {
			name:    "kade",
			number:  10765557221,
			isAdmin: false,
		},
	}

	fmt.Println("Initial users:")
	//usersNameSorted := []string{}
	usersNameSorted := make([]string, 0, len(usersMap))

	for name := range usersMap {
		usersNameSorted = append(usersNameSorted, name)
	}
	for _, name := range usersNameSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	testAdvanceFunc3(usersMap, "john")
	testAdvanceFunc3(usersMap, "santa")
	testAdvanceFunc3(usersMap, "kade")

	fmt.Println("Final users:")
	usersNameSorted = make([]string, 0, len(usersMap))

	for name := range usersMap {
		usersNameSorted = append(usersNameSorted, name)
	}
	for _, name := range usersNameSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	fmt.Println("---")

	// Closures

	// closure is function that references variables
	// from outside its own function body
	// function may access and assign to referenced variables

	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive

	fmt.Println("---")

	// Advanced Function Example 4
	testAdvanceFunc4([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})
	testAdvanceFunc4([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})
	testAdvanceFunc4([]emailBill{
		{743},
		{13},
		{8},
	})

	fmt.Println("---")

	// Anonymous Function
	nums := []int{1, 2, 3, 4, 5}

	// Here we define an anonymous function that doubles an int
	// and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)

	fmt.Println("---")

	// Advanced Function Example 5
	testAdvanceFunc5([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	testAdvanceFunc5([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}
