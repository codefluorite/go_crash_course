package main

import "fmt"

func main() {
	// Define a map (key value pair)

	// emails := make(map[string]string)

	// Assign kv

	// emails["Flora"] = "flora@codefluorite.com"
	// emails["Steph"] = "steph@codefluroite.com"

	// Declare map and add kv

	emails := map[string]string{"Flora": "flora@codefluorite.com", "Steph": "steph@codefluorite.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Flora"])

	// Delete

	delete(emails, "Steph")
	fmt.Println(emails)
}
