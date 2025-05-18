package main

import (
	"fmt"
	"math/rand"
	//"rsc.io/quote"
)

func randomize_int(min, max int) int{
	return rand.Intn(max - min + 1) + min
}

func main() {
	//fmt.Println("Hello go")
	//fmt.Println(quote.Hello())
	//var two int = 2
	//var sixteen int = two << 3
	//fmt.Println(sixteen)
	var random_array [1000]int

	var randomize func(int, int) int = randomize_int

	for index, _ := range random_array{
		random_array[index] = randomize(1, 100) 
		fmt.Print(random_array[index], "\t")
		if (index + 1) % 10 == 0 && index != 0{
			fmt.Println()
		}
	}
	fmt.Println()
}