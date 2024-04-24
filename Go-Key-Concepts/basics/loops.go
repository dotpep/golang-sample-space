package main

import (
	"errors"
	"fmt"
	"time"
)

// Loops
// basic loop
//	for INITiAL; CONDITION; AFTER{
//		// do something
//	}

// forever loop
//	for INITiAL; ...; AFTER{
//		// do something forever
//	}

// while loop
//	for CONDITION {
//		// do some stuff while CONDITION is true
//	}

// Basic loop
func basicLoopExample() {
	numMessage := 4

	totalCost := 0.0
	for i := 0; i < numMessage; i++ {
		iterCost := 0.01 * float64(i)
		totalCost += 1.0 + iterCost
		fmt.Printf("%v message 1.0 + %.2f \n", i+1, iterCost)
	}
	fmt.Printf("total: %.2f\n", totalCost)
}

// Forever loop
func foreverLoopExample(totalCost, thresh float64) float64 {
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return float64(i)
		}
	}
}

func testForeverLoop() {
	result := foreverLoopExample(30, 25)
	fmt.Printf("first total: %.2f\n", result)
	result = foreverLoopExample(15, 25)
	fmt.Printf("second total: %.2f\n", result)
}

// While loop
func forWhileLoopExample(plantHeight, maxPlantHeight int) {
	for plantHeight < maxPlantHeight {
		fmt.Println("still growing! current height: ", plantHeight)
		plantHeight++
		fmt.Println("plant sleep zzzZz")
		time.Sleep(100)
		//time.Sleep(1 * time.Second)
	}
	fmt.Printf("plant has grown to %v inches\n", plantHeight)
}

// Reminder operator %
// 4/3 = 1.33, int(4/3) = 1, 6/3 = 2
// 2%2 == 0 even num, 1%2 == 1 odd num
// 6%3 = 0, 12%4 = 0, 16%5 = 1, 22%8 = 6

// Logical operator 'and'=='&&', 'or'=='||'
// and:
// True && Ture == True
// True && False == False
// or:
// True || Ture == True
// True || False == True
// False || False == False

func threeDigitValidator(num int) (int, error) {
	if num >= 999 {
		return num, errors.New("number is out of 3 digit range")
	}
	if num < 0 {
		return num, errors.New("number is negative")
	}
	return num, nil
}

func intThreeDigitParser(num int) {
	num, err := threeDigitValidator(num)
	if err != nil {
		fmt.Println(err)
		return
	}

	var firstDigit, middleDigit, lastDigit int

	firstDigit = int(num / 100)
	middleDigit = int((num / 10) % 10)
	lastDigit = num % 10

	fmt.Printf(
		"first: %v, second: %v, third: %v digits\n",
		firstDigit, middleDigit, lastDigit,
	)
}

// Continue keyword
func continueKeyLoopExample(count int) {
	for i := 0; i < count; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Break keyword
func breakKeyLoopExample(startNum, stopNum, countLoop, increment int) {
	for i := startNum; i < (startNum + countLoop); i += increment {
		if i == stopNum {
			break
		}
		fmt.Println(i)
	}
}

func main() {
	// Basic loop
	fmt.Println("basic loop:")
	basicLoopExample()

	// Forever loop
	fmt.Println("---")
	fmt.Println("forever loop:")
	testForeverLoop()

	// While loop
	fmt.Println("---")
	fmt.Println("while loop:")
	forWhileLoopExample(0, 20)

	fmt.Println("---")
	// Reminder operator
	intThreeDigitParser(586)
	intThreeDigitParser(10)
	intThreeDigitParser(5)
	intThreeDigitParser(0)
	intThreeDigitParser(-45)
	intThreeDigitParser(58633)

	fmt.Println("---")
	fmt.Println("continue keyword:")
	// Continue keyword
	continueKeyLoopExample(10)

	fmt.Println("---")
	fmt.Println("break keyword:")
	// Break keyword
	breakKeyLoopExample(10, 30, 40, 5)
}
