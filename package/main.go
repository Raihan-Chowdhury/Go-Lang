package main

import (
	"fmt"
	"math"

	"github.com/Projects/package/strutl"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Round(2.3))
	fmt.Println(strutl.Reverse("hello"))
}
