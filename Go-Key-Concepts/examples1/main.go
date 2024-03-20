package main

import (
	"fmt"
	"math"
	"strconv"
)

type student struct {
	name string
	age  int
}

func main() {
	const N = 100
	result := fizzBuzz(N)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

	fmt.Println("---")
	fmt.Println(result)

}

func fizzBuzz(length int) []string {
	slice := make([]string, length)

	for i := 1; i < length+1; i++ {
		iter := i - 1
		var text string

		isInCondition := func(num int) bool {
			if num%3 == 0 || num%5 == 0 {
				return true
			} else {
				return false
			}
		}

		if isInCondition(i) {
			if i%3 == 0 {
				text += "Fizz"
			}
			if i%5 == 0 {
				text += "Buzz"
			}
			slice[iter] = text
		} else {
			slice[iter] = strconv.Itoa(i)
		}
	}

	return slice
}

func ifConditionExample() {
	age := 18
	if age >= 18 {
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are not adult")
	}
	fmt.Println("---")

	loopCount := 100
	for i := 1; i < loopCount; i++ {
		var text string

		if i%3 == 0 {
			text += "Fizz"
		}
		if i%5 == 0 {
			text += "Buzz"
		}
		fmt.Println(i, text)
	}
	fmt.Println("---")
	const SIZE = 15
	slice := make([]string, SIZE)

	//for n := range slice {
	for i := 1; i < SIZE+1; i++ {
		j := i - 1
		var text string

		//var condition bool

		//if i%3 == 0 || i%5 == 0 {
		//	condition = true
		//} else {
		//	condition = false
		//}

		isInCondition := func(num int) bool {
			if num%3 == 0 || num%5 == 0 {
				return true
			} else {
				return false
			}
		}

		if isInCondition(i) {
			if i%3 == 0 {
				text += "Fizz"
			}
			if i%5 == 0 {
				text += "Buzz"
			}
			slice[j] = text
		} else {
			slice[j] = strconv.Itoa(i)
		}

		//switch {
		//case i%3 == 0:
		//	text += "Fizz"
		//case i%5 == 0:
		//	text += "Buzz"
		//case i%3 == 0 && i%5 == 0:
		//	text += "FizzBuzz"
		//default:
		//	text += strconv.Itoa(i)
		//}
		//slice[i-1] = text

		//if i%3 == 0 {
		//	//text += "Fizz"
		//	slice[i] = "Fizz"
		//} else if i%5 == 0 {
		//	//text += "Buzz"
		//	slice[i] = "Buzz"
		//} else if i%3 == 0 && i%5 == 0 {
		//	slice[i] = "FizzBuzz"
		//} else {
		//	slice[i] = strconv.Itoa(i)
		//}
	}
	fmt.Println(slice)

}

func forLoopExample() {
	iterLength := 8

	for i := 1; i < iterLength; i += 2 {
		fmt.Println(i + 1)
	}

	fmt.Println("---")
	for {
		fmt.Println("Hello World!")
		break
	}

	fmt.Println("---")

	const ARRSIZE = 6
	arr := [ARRSIZE]int{1, 2, 3, 4, 5, 6}

	for n := range arr {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n + 1)
	}

	fmt.Println("---")
	i := 1
	for i <= 6 {
		fmt.Println(i)
		i++
	}
}

const GLOBAL_ARRAY_SIZE int = 10

func mathAndConstExample() {
	size := GLOBAL_ARRAY_SIZE
	fmt.Println(size)

	const N = 500000000

	const A = 3e20
	fmt.Println(A)

	const DIVIDED = A / N
	fmt.Println(int64(DIVIDED))

	sineRadian := math.Sin(N)
	roundedRadian := math.Round(sineRadian)
	absoluteRadian := math.Abs(roundedRadian)
	fmt.Println(sineRadian)
	fmt.Println(roundedRadian)
	fmt.Println(absoluteRadian)

	//negativeNum := -1
	var negativeNum float64 = -1.4
	absoluteValue := math.Abs(negativeNum)
	convertedToInt := int(absoluteValue)
	fmt.Println(negativeNum)
	fmt.Println(absoluteValue)
	fmt.Println(convertedToInt)
}

func test() {
	//ternaryOperator := func(boolCondition bool) int {
	//	if boolCondition {
	//		return 1
	//	} else {
	//		return 0
	//	}
	//}
	//fmt.Println(ternaryOperator(false))

	start := 0
	stop := 10
	step := 3
	//compressedSizeMode := true

	//myRange := makeRange(start, stop, step, compressedSizeMode)
	myRange := makeRange(start, stop, step, false)
	//myRange := makeRange(start, stop)
	fmt.Println(myRange)

	//// Call foo with explicit argument
	//result1 := multipleNonProvidedArgs("1 am", false)
	//fmt.Println("Result 1:", result1)

	//// Call foo without argument, uses default value
	//result2 := multipleNonProvidedArgs("any time")
	//fmt.Println("Result 2:", result2)

	// Test cases
	fmt.Println("---")
	fmt.Println("Test cases")
	fmt.Println(makeRange(0, 10))           // [0 1 2 3 4 5 6 7 8 9 10]
	fmt.Println(makeRange(2, 10, 2))        // [2 4 6 8 10]
	fmt.Println(makeRange(0, 10, 3))        // [0 3 6]
	fmt.Println(makeRange(2, 10, false))    // [2 3 4 5 6 7 8 9 10]
	fmt.Println(makeRange(0, 10, true))     // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(makeRange(0, 10, false))    // [0 1 2 3 4 5 6 7 8 9 10]
	fmt.Println(makeRange(0, 10, 3, true))  // [0 3 6]
	fmt.Println(makeRange(0, 10, 3, false)) // [0 3 6 9 12 15 18 21 24 27 30]
	//fmt.Println(makeRange(0, 1000, 2))

	fmt.Println("---")
	//rangeNum, err := makeRange(0, 10)
	//fmt.Println(err)
	rangeNum := makeRange(0, 10)
	fmt.Println(rangeNum)

	//// Test cases with more than 4 optional arguments
	//_, err := makeRange(0, 10, 3, false, true)
	//fmt.Println("Error:", err) // Output: Error: too many optional arguments provided (maximum of 2 allowed)

	//_, err = makeRange(0, 10, 3, false, 1, "string")
	//fmt.Println("Error:", err) // Output: Error: too many optional arguments provided (maximum of 2 allowed)

	// Test cases with more than 4 optional arguments
	//result, err := makeRange(0, 10, 3, false, true)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println(result)
	//}

	//result, err = makeRange(0, 10, 3, false, 1, "string")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println(result)
	//}

	//// Call foo with explicit arguments
	//result1 := foo("1 am", false)
	//fmt.Println("Result 1:", result1)

	//// Call foo with only one argument provided
	//result2 := foo("any time")
	//fmt.Println("Result 2:", result2)

	//// Call foo with explicit arguments
	//result3 := foo("1 pm", false)
	//fmt.Println("Result 3:", result3)

	//// Call foo without any arguments
	//result4 := foo()
	//fmt.Println("Result 4:", result4)
}

func foo(currentTime ...string) string {
	return fooWithMode(true, currentTime...)
}

func fooWithMode(boolMode bool, currentTime ...string) string {
	currentTimeMode := func(time string) string {
		if time == "1 am" {
			return "night"
		}
		return "day"
	}

	modeActivated := "mode activated with: " + currentTimeMode(currentTime[0])
	modeNotActivated := "mode not activated with: " + currentTimeMode(currentTime[0])

	if boolMode {
		return modeActivated
	}
	return modeNotActivated
}

func multipleNonProvidedArgs(currentTime string, boolMode ...bool) string {
	isModeActivated := func(currentMode ...bool) bool {
		if len(currentMode) == 0 {
			return true
		}
		return currentMode[0]
	}

	currentTimeMode := func(currentMode string) string {
		if currentMode == "1 am" {
			return "night"
		} else {
			return "day"
		}
	}

	if isModeActivated(boolMode...) {
		return "mode activated with: " + currentTimeMode(currentTime)
	} else {
		return "mode not activated with: " + currentTimeMode(currentTime)
	}

	//if len(boolMode) == 0 {
	//	return true
	//}
	//return boolMode[0]
}

// makeRange generates a range of integers based on the given start and stop values,
// with an optional step and compressedSizeMode.
//
// Parameters:
//   - start: The starting value of the range (inclusive).
//   - stop: The ending value of the range (inclusive).
//   - optionalArgs: Optional arguments using variadic args, interface.
//   - step (default value: 1) - If provided, the first optional argument (if int) specifies the step size of the range.
//   - compressedSizeMode (default value: true) - If provided, the second optional argument (if bool) specifies the compressedSizeMode,
//     which if true, compresses the range by excluding the last element.
//
// Returns:
//   - An integer slice representing the generated range.
//
// Example usage:
//
//	fmt.Println(makeRange(0, 10))           // [0 1 2 3 4 5 6 7 8 9 10]
//	fmt.Println(makeRange(2, 10, 2))        // [2 4 6 8 10]
//	fmt.Println(makeRange(0, 10, 3))        // [0 3 6]
//	fmt.Println(makeRange(2, 10, false))    // [2 3 4 5 6 7 8 9 10]
//	fmt.Println(makeRange(0, 10, true))     // [0 1 2 3 4 5 6 7 8 9]
//	fmt.Println(makeRange(0, 10, false))    // [0 1 2 3 4 5 6 7 8 9 10]
//	fmt.Println(makeRange(0, 10, 3, true))  // [0 3 6]
//	fmt.Println(makeRange(0, 10, 3, false)) // [0 3 6 9 12 15 18 21 24 27 30]
func makeRange(start int, stop int, optionalArgs ...interface{}) []int {
	// TODO: Implement handling error when passes more than required arguments into optionalArgs as interface (or make in with struct)
	//if len(optionalArgs) > 2 {
	//	return nil, errors.New("too many optional arguments provided")
	//}

	// TODO: Implement start args as non provided arguments and setted by default 0
	//var start int = 0
	var step int = 1
	var compressedSizeMode bool = true

	// Handling optional arguments
	for i := 0; i < len(optionalArgs); i += 2 {
		switch arg := optionalArgs[i].(type) {
		case int:
			step = arg
		case bool:
			compressedSizeMode = arg
			//default:
			//	return nil, fmt.Errorf("unexpected type for optional argument: %d", i+1)
		}

	}

	// Special case handling when both step and compressedSizeMode are provided
	if len(optionalArgs) == 2 {
		compressedSizeMode = optionalArgs[1].(bool)
	}

	// Declaring default values for optional arguments
	isStepProvied := func(number int) bool {
		if number > 0 {
			return true
		} else {
			return false
		}
	}

	isCompressSizeMode := func(defaultMode ...bool) bool {
		if len(defaultMode) == 0 {
			return true
		}
		return defaultMode[0]
	}

	// Logic of making range of numbers in slice
	slice := make([]int, stop-start+1)

	for i := range slice {
		if isStepProvied(step) {
			if isCompressSizeMode(compressedSizeMode) {
				newSize := stop / step
				slice = slice[:newSize]
				calc := start + (i * step)
				slice[i] = calc
				if i == newSize-1 {
					break
				}
			} else {
				slice[i] = start + (i * step)
			}
		} else {
			slice[i] = i + start
		}
	}

	return slice
	//return slice, nil
}

func makeRange2(start int, stop int, step int, compressedSizeMode ...bool) []int {
	slice := make([]int, stop-start+1)

	// TODO: Implement makeRange function can handle 2 args that can be optionally provided when using this function
	// TODO: Implement non provided step argument value and default value for function
	isStepProvied := func(number int) bool {
		if number > 0 {
			return true
		} else {
			return false
		}
	}

	// TODO: Implement function handling CompressMode of Making range of numbers with default is true and can be undefined as arguments of function
	//var compressedSizeMode bool =true
	isCompressSizeMode := func(defaultMode ...bool) bool {
		if len(defaultMode) == 0 {
			return true
		}
		return defaultMode[0]
	}

	for i := range slice {
		//if step > 1 {
		//	if compressedSizeMode && len(slice) > 1 {
		if isStepProvied(step) {
			if isCompressSizeMode(compressedSizeMode...) {
				newSize := stop / step
				slice = slice[:newSize]
				calc := start + (i * step)
				slice[i] = calc
				if i == newSize-1 {
					break
				}
				//} else if !compressedSizeMode {
			} else {
				slice[i] = start + (i * step)
			}
		} else {
			slice[i] = i + start
		}
	}

	return slice
}

func arrayExamples() {
	const ARRAY_SIZE int = 5
	//var arr [ARRAY_SIZE]int
	//fmt.Println("emp:", arr)
	//actuallyLen := len(arr)

	//arr[4] = 100
	//fmt.Println("set:", arr)
	//fmt.Println("get:", arr[4])
	//fmt.Println("len:", actuallyLen)

	filledArray := [ARRAY_SIZE]int{1, 2, 3, 4, 5}
	//fmt.Println("dcl:", filledArray)

	//var twoDimensional [2][3]int
	//fmt.Println("2d:", twoDimensional)

	//fmt.Println("---")
	//fmt.Println("my array before:", arr)

	//arr[1] = 7

	//for i := 0; i < ARRAY_SIZE; i++ {
	//	if arr[i] == 0 {
	//		arr[i] = i + 1
	//	} else if arr[i] >= 100 {
	//		arr[i] *= 2
	//	} else {
	//		arr[i] *= 1000
	//	}
	//}
	//fmt.Println("my array after:", arr)

	fmt.Println("---")
	nestedPrintArray := func(array []int, arraySize int) {
		for i := 2; i < arraySize; i++ {
			fmt.Println(array[i])
		}
	}
	nestedPrintArray(filledArray[:], ARRAY_SIZE)

	fmt.Println("---")
	var studentArray [3]student
	studentArray[0] = student{"John", 19}
	studentArray[1] = student{"Mary", 21}
	studentArray[2] = student{"Peter", 18}

	fmt.Println("students array:", studentArray)
	fmt.Println("---")

	new1 := student{"New 1", 22}
	new2 := student{"New 2", 20}
	new3 := student{"New 3", 21}

	newStudentArray := [...]student{
		new1,
		new2,
		new3,
	}

	fmt.Println("new student array:", newStudentArray)

	fmt.Println("---")
	i := 32
	fmt.Println(&i)

	copyArray := newStudentArray
	fmt.Println("copy arr:", copyArray)
	fmt.Printf("arr address: %p\n", &newStudentArray)
	fmt.Printf("copy arr address: %p\n", &copyArray)
	fmt.Printf("struct instance of new 1 student (%v): %p\n", new1, &new1)
	fmt.Printf("arr address of new 1 student (%v): %p\n", newStudentArray[0], &newStudentArray[0])

}

func sliceExamples() {
	//randomInteger := rand.Int()
	//fmt.Println(randomInteger)

	//fmt.Println("---")
	//const ARRAY_SIZE int = 5
	//var arr [ARRAY_SIZE]int

	//for i := 0; i < ARRAY_SIZE; i++ {
	//	arr[i] = rand.Int()
	//}
	//fmt.Println("arr:", arr)
	//fmt.Println("---")

	start := 0
	stop := 10

	//isStepProvied := false
	//step := 2

	var step int
	//step = 3

	slice := make([]int, stop-start+1)

	//for i := range slice {
	//	if slice[i] != step {
	//		if isStepProvied {
	//			slice[i] = start + (i * step)
	//		} else {
	//			slice[i] = i + start
	//		}
	//	}
	//}

	var compressedSizeMode bool = true
	//isStepProvied := true

	isStepProvied := func(number int) bool {
		if number > 0 {
			return true
		} else {
			return false
		}
	}

	for i := range slice {
		//if step > 0 {
		if isStepProvied(step) {
			if compressedSizeMode {
				newSize := stop / step
				slice = slice[:newSize]
				calc := start + (i * step)
				slice[i] = calc
				//fmt.Printf("iter: %v, newsize: %v, slice: %v, calc: %v \n", i, newSize, slice, calc)
				if i == newSize-1 {
					break
				}
			} else if !compressedSizeMode {
				slice[i] = start + (i * step)
			}
		} else {
			slice[i] = i + start
		}
	}

	fmt.Println(slice)

	//fmt.Println("---")
	//fmt.Println(stop / step)
	//fmt.Println(len(slice))
	////slice = slice[:5]
	//slice = slice[:stop/step]
	//fmt.Println(len(slice))
}
