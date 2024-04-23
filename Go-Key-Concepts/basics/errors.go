package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Handling Errors Example 1
type User struct {
	name string
}

func getUser(user User) (User, error) {
	if user.name == "" {
		return user, fmt.Errorf("User name is not provided and not found")
	}
	return user, nil
}

// Errors Interface Example 2
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	customerCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	spouseCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	//customerCost, customerErr := sendSMS(msgToCustomer)
	//spouseCost, spouseErr := sendSMS(msgToSpouse)
	//if customerErr != nil && spouseErr != nil {
	//	return 0.0, fmt.Errorf("customer: %v, spouse: %v", customerErr, spouseErr)
	//}
	totalCost := customerCost + spouseCost
	return totalCost, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = 0.0002

	msgLength := len(message)

	if msgLength > maxTextLen {
		return 0.0, fmt.Errorf(
			"can't send texts over %v characters because message length: %v",
			maxTextLen, msgLength,
		)
	}
	return costPerChar * float64(msgLength), nil
}

func testErrorHandling1(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("==================")
	fmt.Println("Message for customer: ", msgToCustomer)
	fmt.Println("Message for spouse: ", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

// Errors Formatting Stirngs Review
// we need better error logs for our bachend to debug code easily

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf(
		"SMS that costs $%.2f to be sent to '%s' can not be sent",
		cost, recipient,
	)
}

func testFormattedError2(cost float64, recipient string) {
	sms := getSMSErrorString(cost, recipient)
	fmt.Println(sms)
	fmt.Println("==================")
}

// Custom Error Type with standard 'error interface'
//	type error interface {
//		Error() string
//	}

type userError struct {
	name string
}

func (ue userError) Error() string {
	return fmt.Sprintf("%v has a problem with account", ue.name)
}

func sendSMSValidator(msg, userName string) error {
	if !canSendToUser(userName) {
		return userError{name: userName}
	}
	return nil
}

func canSendToUser(userName string) bool {
	const maxUserNameLen = 50
	if len(userName) > maxUserNameLen {
		return false
	}
	return true
}

// Custom Error Example 3
type divideError struct {
	dividend float64
}

func (diverr divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", diverr.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideError{dividend: dividend}
	}
	return (dividend / divisor), nil
}

func testCustomError3(dividend, divisor float64) {
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	result, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quotient: ", result)
	}
	fmt.Println("==================")
}

// Errors Package
// var err error = errors.New("somthing went wrong")

func divideExample2(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("ERROR: no dividing by 0")
	}
	return (x / y), nil
}

func testPackageErrorHandling4(x, y float64) {
	result, err := divideExample2(x, y)
	fmt.Printf("(%v / %v) result is: %.2f\n", x, y, result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==================")
}

func main() {
	// Handling Errors Example 1
	user, err := getUser(User{
		name: "Alex",
	})
	if err != nil {
		fmt.Println("Failed to get user: ", user, err)
		return // that returns stops program execution process in this line
	}
	fmt.Println("Successfully get user: ", user, err)

	fmt.Println("---")

	// Standard 'error interface'
	// Atoi converts a stringified number to an integer
	intVal, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("couldn't convert: ", err)
		// because "42b" isn't a valid integer, we print:
		// couldn't convert:  strconv.Atoi: parsing "42b": invalid syntax
		// Note:
		// 'parsing "42b"; invalid syntax' is returned by .Error() method of standard 'error interface'
		return
	}
	// if we get here, then
	// intVal was converted successfully
	fmt.Println("Converted integer value: ", intVal)

	fmt.Println("---")

	// Errors Interface Example 2
	testErrorHandling1("Thanks for coming in to our flower shop", "We hope you enjoied your gift")
	testErrorHandling1("Thanks for joining us!", "Have a good day.")

	fmt.Println("---")

	// Errors Formatting Stirngs Review
	testFormattedError2(1.4, "+1 (405) 555 0954")
	testFormattedError2(15.5, "+2 (558) 555 5888")
	testFormattedError2(32.2, "+1 (708) 555 0105")

	fmt.Println("---")

	// Custom Error Example 3
	testCustomError3(10, 0)
	testCustomError3(25, 5)
	testCustomError3(15, 30)
	testCustomError3(6, 0)

	fmt.Println("---")

	// Errors Package
	testPackageErrorHandling4(13, 0)
	testPackageErrorHandling4(18, 4)
}
