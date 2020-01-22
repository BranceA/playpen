package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Did I get Pi? %f\n", math.Pi) 

	var x int = 3

	shortHand := 10

	if shortHand > 10 {
		fmt.Println("Bigger than 10")
	} else if shortHand < 10 {
		fmt.Println("Smaller than 10")
	} else {
		fmt.Println("That is 10")
	}

	fmt.Println(x)

	fmt.Println(shortHand)

	var testArray [5]int
	testArray[2] = 7
	fmt.Println(testArray)

	testArrTwo := [5]int{3, 5, 8, 1, 1}
	fmt.Println(testArrTwo)

	testSlice := []int{2, 3, 4, 5, 6}
	fmt.Println(testSlice)
	testSlice = append(testSlice, 7)
	fmt.Println(testSlice)
}