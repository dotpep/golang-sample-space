package main

import "fmt"

// Generics

type NumericType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sumNumbers[T NumericType](numbers []T) T {
	var result T
	for _, num := range numbers {
		result += num
	}

	return result
}

func main() {
	// numeric data type
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int8{1, 2, 3, 4, 5}
	numbers3 := []int16{1, 2, 3, 4, 5}
	numbers4 := []int32{1, 2, 3, 4, 5}
	numbers5 := []int64{1, 2, 3, 4, 5}
	numbers6 := []float32{1.1, 2.1, 3.1, 4.1, 5.1}
	numbers7 := []float64{1.1, 2.1, 3.1, 4.1, 5.1}

	// function for numeric data type
	fmt.Println(sumNumbers(numbers1))
	fmt.Println(sumNumbers(numbers2))
	fmt.Println(sumNumbers(numbers3))
	fmt.Println(sumNumbers(numbers4))
	fmt.Println(sumNumbers(numbers5))
	fmt.Println(sumNumbers(numbers6))
	fmt.Println(sumNumbers(numbers7))
}

// Duplicative functions that parameter and return value
// is same numeric data type but it is not working with all
// numeric data types, (DRY)
// for that reason we can use Generics

//func sumNumbers(nums []int) int {
//	var result int
//	for _, num := range nums {
//		result += num
//	}

//	return result
//}

//func sumNumbers(nums []float32) float32 {
//	var result float32
//	for _, num := range nums {
//		result += num
//	}

//	return result
//}
