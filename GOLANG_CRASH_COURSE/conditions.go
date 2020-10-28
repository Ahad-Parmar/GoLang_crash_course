package main

import "fmt"

func main() {
	x := 5
	y := 5

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "yellow"
	if color == "yellow" {
		fmt.Println("color is yellow")
	} else if color == "green" {
		fmt.Println("color is green")
	} else {
		fmt.Println("color is neither yellow nor green")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("yellow")
	case "blue":
		fmt.Println("green")
	default:
		fmt.Println("nothing")
	}
}
