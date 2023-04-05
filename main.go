package main

import (
	"fmt"
)

func getDegreesOfTwo(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 1 << uint(i)
	}
	return result
}

func getEvenAndOdd(arr []int) ([]int, []int) {
	even := make([]int, 0)
	odd := make([]int, 0)
	for _, val := range arr {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	return even, odd
}

func sortArray(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func getArithmeticDiff(arr []int) int {
	diff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return 0
		}
	}

	return diff
}

func main() {
	// Exercise 2
	n := 5
	degreesOfTwo := getDegreesOfTwo(n)
	fmt.Println(degreesOfTwo)

	// Exercise 14
	arr := []int{9, 5, 3, 7, 2, 8, 4, 1, 6}
	even, odd := getEvenAndOdd(arr)
	sorted := append(sortArray(even), sortArray(odd)...)
	fmt.Println(sorted)

	// Exercise 21
	arr2 := []int{1, 3, 5, 7, 9}
	diff := getArithmeticDiff(arr2)
	if diff == 0 {
		fmt.Println(diff)
	} else {
		fmt.Printf("The difference of the arithmetic progression is %d\n", diff)
	}
}
