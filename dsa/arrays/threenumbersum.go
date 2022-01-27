// this program takes an array of numbers and finds an array of triplets in which each triplets equals the target sum
package main

import (
	"fmt"
	"sort"
)

func main() {
	answer := ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 1)
	fmt.Printf("%v", answer)
}

// runs in O(n^2) time and O(n) space
func ThreeNumberSum(input []int, targetSum int) [][]int {
	sort.Ints(input) // sort in ascending order first
	triplets := [][]int{}

	// run from 0 to third to the last index, since we're doing a nested loop where we use pointers to the left and right index of the nested
	// loop to determine the potential sums
	for i := 0; i < len(input)-2; i++ {
		left, right := i+1, len(input)-1
		for left < right { // making sure they don't overlap
			currentSum := input[i] + input[left] + input[right]
			if currentSum == targetSum { // if equal simpply add to list containing
				triplets = append(triplets, []int{input[i], input[left], input[right]}) // append
				// move both pointers
				left += 1
				right -= 1
			} else if currentSum < targetSum { // if less, move left closer
				left += 1
			} else { // move right closer
				right -= 1
			}
		}
	}

	return triplets
}
