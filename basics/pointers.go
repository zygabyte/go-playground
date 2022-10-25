package main

import "fmt"

func main() {
	value := 10
	noPointerUpdate(&value)
	fmt.Println(value)

	correctPointerUpdate(&value)
	fmt.Println(value)
}

func noPointerUpdate(val *int) {
	a := 100
	val = &a // giving val a new address won't affect the value variable that was passed in main
	// to this function
}

func correctPointerUpdate(val *int) {
	a := 100
	*val = a // we dereference the pointer, and the assignment puts the new value in the same memory location pointed
	// to by both val in this function and value in main
}
