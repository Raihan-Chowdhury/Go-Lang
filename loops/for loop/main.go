package main

import "fmt"

func main() {
	//Long methode
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	// Short methode
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d\n", i)
	}
	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) { // i%(3*5) == 0
			fmt.Printf("Fizz-Buzz : %d\t", i)
		} else if i%3 == 0 {
			fmt.Printf("Fizz : %d\t", i)
		} else if i%5 == 0 {
			fmt.Printf("Buzz : %d\t", i)
		} else {
			fmt.Printf("Normal : %d\t", i)
		}
	}
}
