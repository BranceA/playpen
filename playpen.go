package main

import (
	"fmt"
	"math/rand"
	"math"
	"errors"
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

	sides := make(map[string]int)

	sides ["Triangle"] = 3
	sides ["Square"] = 4
	sides ["Pentagon"] = 5
	fmt.Println(sides)
	fmt.Println(sides["Square"])
	delete(sides, "Square")
	fmt.Println(sides)

	for i := 0; i < len(testArrTwo); i++ {
		fmt.Println(testArrTwo[i])
	}

	// basically a while loop 
	j := 0
	for j < len(testArray) {
		fmt.Println(testArray[j])
		j ++
	}

	// space between each comma separated value in println 
	for index, value := range testSlice{
		fmt.Println("index:", index, "value:", value)
	}

	for key, value := range sides{
		fmt.Println("key:", key, "value", value)
	}

	fmt.Println(sum(8, 11))

	result, err := sqrt(15)

	// function always returns a value for err and it is nil if there are no errors 
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	testPerson := person{name: "Brance", age: 31}
	fmt.Println(testPerson)
	fmt.Println(testPerson.name)

	thisIsCool := 4
	inc(&thisIsCool)
	fmt.Println(thisIsCool)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	// the returned error value is nil if we make it this far
	return math.Sqrt(x), nil
}

type person struct {
	name string
	age int
}

func inc(x *int){
	*x++
}