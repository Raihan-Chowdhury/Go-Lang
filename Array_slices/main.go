package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	// Declare and assign
	MyFruit := [2]string{"Apple", "Orange"}

	fruitSlicer := []string{"A", "B", "C"}

	fmt.Println(fruitArr)
	fmt.Println(MyFruit)
	fmt.Println(fruitSlicer[1:2])
}
