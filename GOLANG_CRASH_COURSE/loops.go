package main

import "fmt"

func main() {
	// Long method
	a := 1
	for a <= 10 {
		fmt.Println(a)
		// a+=1
		a++
	}
	// short method
	for a := 1; a <= 10; a++ {
		fmt.Printf("Number %d\n", a)
	}

	// FizzBuzz challenge
	for a := 1; a <= 100; a++ {
		if a%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if a%3 == 0 {
			fmt.Println("Fizz")
		} else if a%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(a)
		}
	}
}
