package main

import (
	"errors"
	"fmt"
	"reflect"
)

func printSliceType(array interface{}) {
	switch val := array.(type) {
	case []int:
		fmt.Printf("integer slice: %v\n", val)
	case []string:
		fmt.Printf("string slice: %v\n", val)
	default:
		fmt.Printf("default error\n")
	}
}

func printArray(arr []int, arrSize int) {
	for i := 0; i < arrSize; i++ {
		fmt.Printf("iter: %v, val: %v\n", i+1, arr[i])
	}
}

// Slices Example 1
const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMsgArr := getMessageWithRetries()
	if plan == planPro {
		return allMsgArr[:], nil // copy array to slice
	}
	if plan == planFree {
		return allMsgArr[0:2], nil // copy arr to slice with index range
	}
	return nil, errors.New("unsupported plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sigh up",
		"pretty please click here",
		"we beg you to sigh up",
	}
}

func testSlice1(name string, doneAt int, plan string) {
	defer fmt.Println("==================")
	fmt.Printf("sending to %v...\n", name)

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("sending: '%s'\n", msg)
		if i == doneAt {
			fmt.Println("they responded")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

// Slices Example 2
func getMessageCosts(messages []string) []float64 {
	lengthOfMsgs := len(messages)
	costs := make([]float64, lengthOfMsgs)
	for i := 0; i < lengthOfMsgs; i++ {
		msg := messages[i]
		cost := float64(len(msg)) * 0.01
		costs[i] = cost
	}
	return costs
}

func testAlocatedSlice2(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages: ")
	printSlice(messages)
	fmt.Println("Costs: ")
	printSlice(costs)
	fmt.Println("==================")
}

func printSlice(slice interface{}) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fmt.Println("- ", val.Index(i))
		}
	default:
		fmt.Println("ERROR: not a slice passed")
	}
}

// Variadic function and Spread operator
// Slice Example 3
func sum(nums ...float64) float64 {
	var total float64
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		total += num
	}
	return total
}

func testVariadicAndSpreadSlice3(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("==================")
}

func variadicFunc(sliceParam ...int) {
	fmt.Println("our slice of int: ", sliceParam)
}

// Slice Example 4
// append build-in func
type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	//costsByDay := []float64{}

	//for i := 0; i < len(costs); i++ {
	//	cost := costs[i]
	//	for cost.day >= len(costsByDay) {
	//		costsByDay = append(costsByDay, 0.0)
	//	}
	//	costsByDay[cost.day] = cost.value
	//}

	// find maximum of day value
	var maxDay int = 0
	for _, cost := range costs {
		if cost.day > maxDay {
			maxDay = cost.day
		}
	}

	// pre allocate slice with maximum length
	costsByDay := make([]float64, maxDay+1)

	// assign cost values by day
	for _, cost := range costs {
		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}

func testAppendSlice4(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("==================")
}

// 2d Slice Example 5
func createMatrix(rows, cols int) [][]int {
	//matrix := make([][]int, 0)

	//for i := 0; i < rows; i++ {
	//	row := make([]int, 0)
	//	for j := 0; j < cols; j++ {
	//		row = append(row, i*j)
	//	}
	//	matrix = append(matrix, row)
	//}

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func test2DMatrixSlice5(rows, cols int) {
	fmt.Printf("Creating %vx%v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("==================")
}

// Range Slices Example 6
func indexOfFirstBadWord(wordsMsg []string, badWords []string) int {
	for i, word := range wordsMsg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func testRangeSlices6(wordsMsg []string, badWords []string) {
	i := indexOfFirstBadWord(wordsMsg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", wordsMsg)
	for _, badWord := range badWords {
		fmt.Println(" -", badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("==================")
}

func main() {
	// Slice with fixed size in brackets
	// defined in compile time
	// is Array not a slice
	const SIZE = 6
	myArray := [SIZE]int{4, 2, 3, 5, 6, 1}
	fmt.Println("Array in golang: ", myArray)

	// Slices are dynamic and flexible
	// because they do not have a fixed Size
	// at compile time and refers to array structure
	sliceInt := []int{1, 2, 3, 4}
	sliceStr := []string{"a", "b", "c", "d"}
	printSliceType(sliceInt)
	printSliceType(sliceStr)

	fmt.Println("---")
	printArray(myArray[:], SIZE)

	fmt.Println("---")
	low := 1
	high := 5 // high-1
	sliceCopyArray := myArray[low:high]
	fmt.Println(sliceCopyArray)

	fmt.Println("---")
	// arrayName[lowIndex:highIndex]
	// arrayName[lowIndex:] // cut with left side
	// arrayName[:highIndex] // cut with right side
	// arrayName[:]  // copy entire array to slice
	testSlice1("Bob", 3, planFree)
	testSlice1("Rob", 3, planPro)
	testSlice1("Sally", 2, planPro)
	testSlice1("Jeff", 3, "no plan")

	fmt.Println("---")

	// Make function
	// that create slice with Size and Capacity in compile time
	// with 0 values of length and capacity in alocated RAM memory Address:Data[length]
	// like array with fixed size
	// this make slice works underhood of created fixed array
	// for perfomance reasons
	// make it used for perfomance of slice
	// func make([]T, len, cap)
	sizeLengthSlice := 5
	capacitySlice := 10
	alocatedSlice := make([]int, sizeLengthSlice, capacitySlice)
	alocatedSlice2 := make([]int, 4)
	fmt.Printf(
		"slice with alocated size and capacity: %v,\nslice with only size: %v\n",
		alocatedSlice, alocatedSlice2,
	)
	// build-in Length and Capacity function
	// to get size and capacity of array, slice
	// make function:
	// Size is create slice with 0 values filled into it
	// Capacity is max size that alocated for that slice
	// referenced to array that fixed size of Capacity
	fmt.Println("length size of alocated slice: ", len(alocatedSlice))
	fmt.Println("array underhood capacity of alocated slice: ", cap(alocatedSlice))

	fmt.Println("---")

	// Slices Example 2
	msgSlice1 := []string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	}
	msgSlice2 := []string{
		"I don't want to be here anymore",
		"Can we go home",
		"I'm hungry",
		"I'm bored",
	}
	testAlocatedSlice2(msgSlice1)
	testAlocatedSlice2(msgSlice2)

	fmt.Println("---")

	// Variadic function
	// defines that we can fill func parameter
	// argument like fill slices of integer
	// that would be stored in func param as slice of int
	// is like *args in python
	//func sum(nums ...int) int {
	//	// nums is just a slice
	//	for i := 0; i < len(nums); i++ {
	//		num := nums[i]
	//	}
	//}
	//func main(){
	//	total := sum(1, 2, 3)
	//	fmt.Println(total)
	//}

	// Spread operator
	// is like inversion of Variadic func
	// but for giving arguments as slice for variadic func
	// used to slice our array slices
	// and give this as arguments for variaduc func
	// is like my_func(*array) in python to give arg in sliced form array to func param
	spreadOperatorSliceArgs := []int{1, 2, 3, 4, 5}
	variadicFunc(spreadOperatorSliceArgs...)

	// Variadic function and Spread operator
	// Slice Example 3
	testVariadicAndSpreadSlice3(2, 4, 4.2, 5.5)
	testVariadicAndSpreadSlice3(12.9, 5.4)
	testVariadicAndSpreadSlice3(19, 2, 3.99)

	fmt.Println("---")

	// Append function
	// append is Variadic following valid:
	// slice = append(slice, oneThing)
	// slice = append(slice, firstThing, secondThing)
	// slice = append(slice, anotherSlice...)

	// dont do this!:
	// someSlice = append(otherSlice, element)

	// Slice Example 4
	// append build-in func
	testAppendSlice4([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	testAppendSlice4([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})

	fmt.Println("---")

	// Slice of slices
	// 2d matrix
	// rows := [][]int{}

	// 2d Slice Example 5
	test2DMatrixSlice5(2, 2)
	test2DMatrixSlice5(10, 10)

	fmt.Println("---")

	// Range for loop for Slices
	// is like foreach loop in c# but in golang
	// it uses for slice, array and etc data structe
	// that can be iterable
	// and we can retrieve index and value of this slice of fruitList
	// for INDEX, ELEMENT_VALUE := range SLICE {}
	fruitList := []string{"apple", "banana", "grape", "kiwi"}
	for i, fruit := range fruitList {
		fmt.Printf("index: %v, fruit: %v\n", i, fruit)
	}

	fmt.Println("---")

	// Range Slices Example 6
	badWords := []string{"crap", "shoot", "dang", "frick"}
	wordsMessage := []string{"hey", "there", "john"}
	testRangeSlices6(wordsMessage, badWords)

	wordsMessage = []string{"ugh", "oh", "my", "frick"}
	testRangeSlices6(wordsMessage, badWords)
}
