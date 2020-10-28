package main

import "fmt"

func main() {
	i := 10
	j := &i

	fmt.Println(i, j)
	fmt.Printf("%T\n", j)

	// Use * to read val from address
	fmt.Println(*j)
	// or
	fmt.Println(*&i)

	// Change val with pointer
	*j = 15
	fmt.Println(i)
}
