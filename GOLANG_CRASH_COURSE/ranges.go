package main

import "fmt"

func main() {
	ids := []int{23, 52, 66, 22, 10, 84}

	// Loop through ids using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "Bob@gmail.com",
		"joy": "joy@gmail.com",
		"marry":    "marry@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name:", k)
	}
}
