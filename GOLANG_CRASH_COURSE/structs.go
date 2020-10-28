package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	// fName  string
	// lName  string
	// city   string
	// gender string
	// age    int

	fName, lName, city, gender string
	age                        int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, I am  " + p.fName + " " + p.lName + " I am " + strconv.Itoa(p.age) +" years old "
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lName = spouseLastName
	}
}
func main() {
	// Init person using struct
	person1 := Person{fName: "ahad", lName: "parmar", city: "amreli", gender: "M", age: 21}
	person2 := Person{fName: "ayesha", lName: "parmar", city: "surat", gender: "F", age: 26}

	// Alternative
	// person1 := Person{"ahad", "parmar", "amreli", "M", 21}
	// fmt.Println(person1)

	// fmt.Println(person1.fName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()

	person2.getMarried("parmar")
	fmt.Println(person2.greet())

	person1.getMarried("No one")
	fmt.Println(person1.greet())
	// fmt.Println(person2)
}
