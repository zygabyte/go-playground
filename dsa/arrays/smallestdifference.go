// this program takes to arrays and finds two numbers from each of the array, which have the smallest possible difference (wrt 0)
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	input1 := []int{-1, 5, 10, 20, 28, 3}
	input2 := []int{26, 134, 135, 15, 17}
	answer := SmallestDifference(input1, input2)
	fmt.Printf("%v", answer)
}

// runs in O(nlogn) + O(mlogm) time and O(1)
func SmallestDifference(input1, input2 []int) []int {
	sort.Ints(input1)
	sort.Ints(input2)
	idx1, idx2 := 0, 0 // use two pointers for the two arrays
	smallest, current := math.MaxInt32, math.MaxInt32
	smallerPair := []int{}

	// idea is this, we take two numbers from the array (starting with the first in each). then we check which is greater than the other
	// if the first is larger, then we want to move the pointer in the first to close up the space, and vice versa. if they are same, we've found the pair
	// so for each move, we find the current, and see difference between the two (absolute), and then see if that is less than the current smallest difference
	// if so, then current is the new smallest
	for idx1 < len(input1) && idx2 < len(input2) {
		first, second := input1[idx1], input2[idx2]
		current = int(math.Abs(float64(first) - float64(second)))
		if first < second {
			idx1 += 1
		} else if second < first {
			idx2 += 1
		} else {
			return []int{first, second}
		}

		if current < smallest {
			smallest = current
			smallerPair = []int{first, second}
		}
	}

	return smallerPair
}
