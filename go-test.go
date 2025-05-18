package main

import (
	"fmt"
	"math/rand"
)

func action(n1 int, n2 int, operation func(int, int) int) int {
	return operation(n1, n2)
}

func add(x int, y int) int      { return x + y }
func multiply(x int, y int) int { return x * y }

func display(message string) {
	fmt.Println(message)
}

func main() {
	/*
		fmt.Print("Input the length of array\n")
		var array_length int
		fmt.Scanln(&array_length)
		var array, length = do_something(array_length)
		for index, value := range array {
			fmt.Println(index, "\t", value, "\n")
		}
		fmt.Println(length)

		f := add             //или так var f func(int, int) int = add
		fmt.Println(f(3, 4)) // 7

		f = multiply         // теперь переменная f указывает на функцию multiply
		fmt.Println(f(3, 4)) // 12

		// f = display      // ошибка, так как функция display имеет тип func(string)

		var f1 func(string) = display // норм
		f1("hello")

		fmt.Println(action(10, 25, add)) // 35
		action(5, 6, multiply)           // 30

		users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
		// удаляем 4-й элемент
		var n = 3
		users = append(users[:n], users[n+1:]...)
		fmt.Println(users) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

		var message string = "Hello, here goes your string ^w^"
		hello_world(message, func(mess string) string { return mess + "\nI added a new line\n" })
	*/
	var min, max int = 5, 100
	var list = new(List)
	for i := 0; i < 100; i++ {
		if i == 0 {
			list.create(rand.Intn(max-min+1) + min)
		} else {
			list.add(rand.Intn(max-min+1) + min)
		}
	}
	list.print()
}

func do_something(array_length int) ([]uint8, int) {
	var array []uint8 = make([]uint8, int(array_length))
	var i int = 0
	for ; i < array_length; i++ {
		array[i] = uint8(rand.Uint32() % 256)
		fmt.Print(array[i], " ")
	}

	fmt.Println()

	for i := 0; i < array_length; i++ {
		for j := 0; j < array_length; j++ {
			fmt.Print(int(array[i])*int(array[j]), "\t")
		}
		fmt.Println()
	}

	fmt.Println()

	var users []string = []string{"Tom", "Bob", "Alice"}
	for index, value := range users {
		fmt.Print(index, value, "\n")
	}

	fmt.Println()

	length := len(array)

	return array, length
}

func hello_world(message string, operation func(string) string) {
	var result string = operation(message)
	fmt.Println(result)
}

type person struct {
	name string
	age  int
}
