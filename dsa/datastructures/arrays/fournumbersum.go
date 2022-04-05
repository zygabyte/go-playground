// this program extends the three number sum. it takes an array of numbers and finds an array of quaduplets in which each quadruplets equals the target sum
package main

import "fmt"

func main() {
	// input := []int{7, 6, 4, -1, 1, 2}
	// target := 16
	input := []int{1, 2, 3, 4, 5, -5, 6, -6}
	target := 5
	fmt.Printf("%v", FourNumberSum(input, target))
}

// O(n^2) time | O(n^2) space - average case
// O(n^3) time | O(n^2) space - worst case
func FourNumberSum(array []int, target int) [][]int {
	// Write your code here.
	allPairSums := map[int][][]int{}    // map to store mapping of already checked numbers (key), and the array of arrays summing that number
	quadruplets := [][]int{}            // final answer to be returned
	for i := 1; i < len(array)-1; i++ { // loop from first element to the second to the last. because we only add a number to the map by the second time we're getting to the number
		// this prevents a duplicate entry
		for j := i + 1; j < len(array); j++ { // inner loop is from the next from the current
			currentSum := array[i] + array[j]
			difference := target - currentSum
			if pairs, ok := allPairSums[difference]; ok { // check if the difference is currently in the map
				for _, pair := range pairs { // if it's there, we wanna loop through all the values (array of arrays making the sum).
					quad := append(pair, array[i], array[j]) //  and then for each of them, we take it's value and with the current value of array[i] and array[j]
					quadruplets = append(quadruplets, quad)  // that makes a quadruplets
				}
			}
		}

		for k := 0; k < i; k++ { // in this loop is where we add the values to the map (not the first time we get them)
			currentSum := array[k] + array[i]
			allPairSums[currentSum] = append(allPairSums[currentSum], []int{array[k], array[i]}) //key is the sum, value is the array itself
		}
	}
	return quadruplets
}
