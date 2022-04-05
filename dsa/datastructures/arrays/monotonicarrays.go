// a monotonic array is an array that is either non decreasing or non increasing.
// i.e. the if trending down (non increasing), it stays that way all through, while if trending up (non decreasing), it stays that way all through
package main

import "fmt"

func main() {
	input := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Printf("%v", IsMonotonic(input))
}

// O (n) time and O (1) space
func IsMonotonic(array []int) bool {
	if len(array) <= 2 { // has to be monotonic if length of 2 or less
		return true
	}

	var i int
	var isIncreasing bool
	// so we try to determine the direction first. afterwhich we can try negating that, which if true, then it isn't monotonic
	for i = 1; i < len(array); i++ {
		if array[i] == array[i-1] { // test for beginning parts of array. if true, numbers are same, so continue till we find direction
			continue
		}
		if array[i] > array[i-1] {
			isIncreasing = true
			break
		} else {
			isIncreasing = false
			break
		}
	}

	for i += 1; i < len(array); i++ {
		if (isIncreasing && array[i] < array[i-1]) || (!isIncreasing && array[i] > array[i-1]) { // try negating the direction, which if it proves true, then it's monotonic
			return false
		}
	}
	return true
}

// O (n) time and O (1) space
// func IsMonotonic(array []int) bool {
// 	// assumption here is we assume that it's both non decreasing and non increasing at the same time. and try to prove otherwise
// 	isNonDecreasing := true // trending up
// 	isNonIncreasing := true // trending down

// 	for i := 1; i < len(array); i++ {
// 		if array[i] < array[i-1] {
// 			isNonDecreasing = false // we're trending down
// 		}
// 		if array[i] > array[i-1] {
// 			isNonIncreasing = false // we're trending up
// 		}
// 	}

// 	return isNonDecreasing || isNonIncreasing
// }
