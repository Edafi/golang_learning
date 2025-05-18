package main

import (
	"fmt"
	"sync"
)

func factorial(number int, int_chanel chan int, go_group *sync.WaitGroup) {
	defer go_group.Done()
	var result int = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	int_chanel <- number
	int_chanel <- result
}

func main() {
	var go_routine_group sync.WaitGroup
	var max_value int = 10
	var int_chanel chan int = make(chan int, max_value*2) // канал для int
	go_routine_group.Add(max_value)
	for i := 1; i <= max_value; i++ {
		go factorial(i, int_chanel, &go_routine_group)
	}

	go func() {
		go_routine_group.Wait()
		close(int_chanel)
	}()

	for {
		number, open_chan1 := <-int_chanel
		factorial, open_chan2 := <-int_chanel

		if !(open_chan1 || open_chan2) {
			fmt.Println("Chanel is closed")
			break
		} else {
			fmt.Println(number, "\t - \t", factorial)
		}
	}
}
