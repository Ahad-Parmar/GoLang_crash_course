package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	//Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["joy"] = "joy@gmail.com"
	// emails["marry"] = "marry@gmail.com"

	// DEclare map and add kv
	emails := map[string]string{"Bob": "Bob@gmail.com",
		"joy": "joy@gmail.com",
		"marry":    "marry@gmail.com"}

	emails["ahad"] = "ahad@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["joy"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
