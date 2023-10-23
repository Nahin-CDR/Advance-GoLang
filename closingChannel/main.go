package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 12

	val, ok := <-ch
	fmt.Println(val, ok)
	/// closing channel
	close(ch)

	val, ok = <-ch
	fmt.Println(val, ok)

	/// channel is closed and you got
	/// false as ok value
	val, ok = <-ch
	fmt.Println(val, ok)

	/// what is panic situation ?
	/*
		    it means on runtime
			when a channel is closed but you
			want to send data to that channel
			or a channel is closed but you
	        want to close it again ..

	*/
}
