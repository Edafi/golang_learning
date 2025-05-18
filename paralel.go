package main

import "fmt"

func main() {

	var int_chanal chan int = make(chan int, 2)
	for i := 1; i <= 5; i++ {
		go factorial(i, int_chanal)
	}
	for {
		number, open_chan_1 := <-int_chanal
		factorial, open_chan_2 := <-int_chanal
		if !(open_chan_1 && open_chan_2) {
			break
		}
		fmt.Println(number, "\t --- \t", factorial)
	}
}

func factorial(n int, chanal_int chan int) {
	//defer close(chanal_int)
	var factorial int = 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	chanal_int <- n
	chanal_int <- factorial
}
