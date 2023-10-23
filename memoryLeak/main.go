package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}

func leak(wg *sync.WaitGroup) {
	ch := make(chan string, 1) // Use a buffered channel with a capacity of 1
	//ch <- "aaaa"
	//close(ch)
	go func() {
		val := <-ch
		fmt.Println("Received:", val)
		wg.Done()
	}()

	fmt.Println("Exiting method")
	wg.Done()
}

/*
If you don't specify a capacity for the channel
when creating it (i.e., you use ch := make(chan string)),
you'll create an unbuffered channel. In this case,
the code will likely still deadlock,
because the sending of the value ch <- "aaaa" in the
leak function will block until there is a corresponding
receiver for the value.

With an unbuffered channel,
a sender and a receiver must be ready at the same
time for communication to occur. In your code,
you have a goroutine that sends a value into the channel,
but there is no goroutine actively waiting to receive the
value from the channel.
This leads to a deadlock situation because the sender
is blocked, waiting for a receiver, and there is no receiver.

To avoid the deadlock, you can either use a buffered
channel as demonstrated in the previous code or ensure
that there is a goroutine actively waiting to receive the
value before sending it into the unbuffered channel.
*/
