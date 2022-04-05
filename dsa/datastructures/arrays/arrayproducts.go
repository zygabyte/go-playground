// we take a non array and get the product of all the elements in the array except the current index in question.
// e.g. [5, 1, 4, 2] becomes [8, 40, 10, 20]
package main

import "fmt"

func main() {
	array := []int{5, 1, 4, 2}
	fmt.Printf("%v", ArrayProducts(array))
}

// O(n) time and O(n) space
func ArrayProducts(array []int) []int {
	products := make([]int, len(array))
	// start every item as 1 in the array
	for i := range array {
		products[i] = 1
	}

	leftProd, rightProd := 1, 1
	// we loop to get the cummulative products on the left, so as we progress, by the end of the array we have all the products to the left of each index
	for i := range array {
		products[i] = leftProd
		leftProd *= array[i]
	}

	// we loop to get the cummulative products on the right, starting from the end of the array, so as we progress, by the end of the array we have all the products to the right of each index
	// this is done to prevent a loop per each index to get the left and right (another solution) . but that'll run in O(n^2) time.
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightProd
		rightProd *= array[i]
	}

	return products
}
