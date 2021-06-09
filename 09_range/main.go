package main

import "fmt"

func main() {
	ids := []int{15, 23, 34, 56, 77, 80}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// NOt using index - put underscore if not in use
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map

	emails := map[string]string{"Flora": "flora@codefluorite.com", "Steph": "steph@codefluorite.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
