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

	shortHand := 5

	fmt.Println(x)

	fmt.Println(shortHand)
}