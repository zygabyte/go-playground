// this program illustrates traversing through a two dimensional array in a zigzag manner.
package main

import "fmt"

func main() {
	array := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16}}

	result := fmt.Sprintf("%v", ZigzagTraverse(array))
	fmt.Print(result)

}

func ZigzagTraverse(array [][]int) []int {
	height := len(array) - 1
	width := len(array[0]) - 1
	result := []int{} // storing the final one d array
	row, col := 0, 0
	goingDown := true // we start based on the information that the first direction is going down

	for !isOutOfBounds(row, col, height, width) { // as long as it's not out of bounds
		result = append(result, array[row][col]) // it's here we append the value in row, col in the 2 d array
		if goingDown {                           // if going down we wanna ensure we aren't past the edges (last row down and last col left)
			if col == 0 || row == height {
				goingDown = false  // wanna change direction now to going up since we're at the edges
				if row == height { // check for this first, so we don't attempt increasing row to out of bounds, which will happen if you put the second check below as the first
					col++ // move to the right
				} else { // at the beginning where col is 0, move down
					row++
				}
			} else { // if not, keep moving downwards.
				row++
				col--
			}
		} else { // moving upwards, we wanna ensure we aren't past the edges (first row up and last col right)
			if row == 0 || col == width { //
				goingDown = true  // wanna change direction now to going down since we're at the edges
				if col == width { // check for this first, so we don't attempt increasing col to out of bounds, which will happen if you put the second check below as the first
					row++ // move down
				} else { // at the beginning where row is 0, move left
					col++
				}
			} else { // keep moving upwards
				row--
				col++
			}
		}
	}
	return result
}

func isOutOfBounds(row, col, height, width int) bool {
	return row < 0 || row > height || col < 0 || col > width
}
