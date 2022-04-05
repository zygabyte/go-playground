package main

import "fmt"

func main() {
	ch := make(chan int) // creates a channel called ch
	a := <-ch            // reads a value from channel into a. it's an unbuffered channel
	b := 300
	ch <- b // reads value of b into ch

	var ch2 <-chan int = make(chan int) //readonly channel
	var ch3 chan<- int = make(chan int) //writeonly channel

	close(ch)
	close(ch3)

	v, ok := <-ch // reading from a closed channel will make ok false
	if !ok {
		fmt.Println("channel is closed")
	}

	bufferCh := make(chan int, 10) // buffered channel
	var bufferCh2 chan int
	bufferCh2 = make(chan int, 10) // buffered channel
}
