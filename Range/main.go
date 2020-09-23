package main

import "fmt"

func main() {
	ids := []int{11, 90, 22, 44, 55}

	// Loop through ids
	for index, id := range ids {
		fmt.Printf("%d	%d\n", index, id)
	}
	// Only value
	for _, id := range ids {
		fmt.Printf("ID : %d\t", id)
	}

	fmt.Println("")

	// loop through map
	web := map[string]string{"bob": "bob.com", "ros": "ros.com"}
	for k, v := range web {
		fmt.Printf("%s	%s\n", k, v)
	}
}
