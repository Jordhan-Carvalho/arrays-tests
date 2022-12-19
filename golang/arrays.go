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
    fmt.Printf("%v %v \n", value, elapsed)
	}

	println("Testing push")
	for _, value := range tests {
		populateArray(value)
		elapsed := timeElapsed(push, value-1)
    fmt.Printf("%v %v \n", value, elapsed)
	}


  // WOW EVEN WORSE THAN JS, maybe its you shit implementation?
	println("Testing unshift")
	for _, value := range tests {
		populateArray(value)
		elapsed := timeElapsed(unshift, value-1)
    fmt.Printf("%v %v \n", value, elapsed)
	}
}

func populateArray(number int) {
	arrayToTest = []int{}
	for i := 0; i < number; i++ {
		arrayToTest = append(arrayToTest, i)
	}
}

type FunctionType = func(int)
func timeElapsed(f FunctionType, arg int) int {
	start := time.Now()
	f(arg)
	elapsed := time.Since(start).Milliseconds()
	return int(elapsed)
}

func get(index int) {
  // The a is here to remove the anoying "declared but not used error"
  a := arrayToTest[index]
  _ = a
}

func unshift(number int) {
  for i:=0; i<number; i++ {
   arrayToTest = append([]int{9}, arrayToTest...) 
  }
}

func push(number int) {
  for i:=0; i<number; i++ {
    arrayToTest = append(arrayToTest, 8)
  }
}
