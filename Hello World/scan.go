package main

import "fmt"

func main() {
	fmt.Println("Enter Your First Name: ")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print("Your Full Name is: ")
	fmt.Print(first + " " + second)
}
