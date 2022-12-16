package main

import (
	"fmt"
	"time"
)

var arrayToTest []int

func main() {
	fmt.Println("hi")
	tests := [...]int{10, 100, 1000, 10000, 100000, 1000000, 10000000}

	println("Testing get")
	for _, value := range tests {
		populateArray(value)
		elapsed := timeElapsed(get, value-1)
		fmt.Println("Time elapsed: AAA", elapsed)
	}


	println("Testing shift")
	for _, value := range tests {
		populateArray(value)
		elapsed := timeElapsed(unshift, value-1)
		fmt.Println("Time elapsed: AAA", elapsed)
	}
}

func populateArray(number int) {
	arrayToTest = []int{}
	fmt.Println("Reseted array capacity: ", cap(arrayToTest))
	for i := 0; i < number; i++ {
		arrayToTest = append(arrayToTest, i)
	}
}

type FunctionType = func(int) int
func timeElapsed(f FunctionType, arg int) int {
	start := time.Now()
	f(arg)
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("Time elapsed: ", elapsed)
	return int(elapsed)
}

func get(index int) int{
		return arrayToTest[index]
}

func unshift(number int) {
  for i:=0; i<number; i++ {

  }
}
