package main

import "fmt"

func main() {
	//MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte af uint8
	// rune af int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name = "Raihan Chowdhury" // var name  string = "Brad"
	var age = 37                  //var age int = 37
	var AreYouSick = false        //var AreYouSick bool = false

	// Shorthand
	University := "DCC"
	// Shorthand With Multiple variable
	username, password := "raihan", "Zushi"

	fmt.Println(name, age, AreYouSick, University)
	fmt.Printf("%T %T %T %T\n", name, age, AreYouSick, University)
	fmt.Println(username, password)
}
