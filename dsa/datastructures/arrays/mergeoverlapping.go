// this program merges intervals of arrays that overlap
// so the input is a two dimensional array where we attempt to determine which of the inner array overlaps.
// for example, an array input of [[1 2] [3 5] [4 7] [6 8] [9 10]] will yield an output of [[1 2] [3 8] [9 10]]
package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}}
	fmt.Printf("%v", MergeOverlappingIntervals(intervals))
}

// O(nlogn) time | O(n) space
func MergeOverlappingIntervals(intervals [][]int) [][]int {
	// Write your code here.
	sortedIntervals := make([][]int, len(intervals))
	va := copy(sortedIntervals, intervals) // copies the source to the destination, which will be sorted based on the first element in each of the inner array
	fmt.Printf("size -> %v \n", va)
	sort.Slice(sortedIntervals, func(a, b int) bool {
		return sortedIntervals[a][0] < sortedIntervals[b][0] // sorting happens here
	})

	currentInterval := sortedIntervals[0]                      // so we take the first element as the first interval and attempt to see if the subsequent elements can fall within this interval
	var mergedIntervals [][]int                                // final output merged overlap
	mergedIntervals = append(mergedIntervals, currentInterval) // append the current interval. as the first interval

	for i := 1; i < len(sortedIntervals); i++ { // a micro optimization here... we loop from the next item (instead of the first, which is already the first)..
		currentEndInterval := currentInterval[1]                                           // take the end bounds of current interval
		nextStartInterval, nextEndInterval := sortedIntervals[i][0], sortedIntervals[i][1] // take start and end bounds of next sorted intervals

		if currentEndInterval >= nextStartInterval { // meaning the next interval belongs within the current interval, hence we can expand the current interval
			currentInterval[1] = max(currentEndInterval, nextEndInterval)
		} else { // new interval. hence append to the merged intervals
			currentInterval = sortedIntervals[i]
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
