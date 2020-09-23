package main

import "fmt"

func main() {
	x := 5
	if x < 10 {
		fmt.Printf("%d is less than 10", x)
		fmt.Println("")
	} else if x == 10 {
		fmt.Printf("%d is equal to 10", x)
		fmt.Println("")
	} else {
		fmt.Printf("%d is greater than 10", x)
		fmt.Println("")
	}

	color := "red"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}

}
