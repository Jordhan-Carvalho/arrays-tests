// https://go.dev/blog/slices-intro
/* package main

import (
	"fmt"
	"time"
)

func main() {
	testingAppendAndPreprendBigO()
	// testingCapacity()

	// testingStuff()
}

func testingStuff() {
	var x []int
	for i := 2; i < 10; i += 2 {
		x = append(x, i)
		// Will doing this the capacity will be 4 ??
	}
	fmt.Println("Testing stuff cap is:", cap(x))
}

func testingCapacity() {
	x := []int{2, 4, 5, 8}

	fmt.Println("The capacity of x is: ", cap(x))

	fmt.Println(x)

	x = append(x, 9)

	fmt.Println("After apend a number the cap is: ", cap(x))
	// The aboce append will double the capctiy

	x = append([]int{1}, x...)

	fmt.Println("After prepending the cap is: ", cap(x))
	// DA FUK... after prepending the cap is 6
	// WOULD THAT MENAS THAT PREPENDING IS SLOWER? SINCE IT DOESNT HAVE EXTRA CAPACITY ???

	fmt.Println(x)

}

// WOW prepending is WAY slower
func testingAppendAndPreprendBigO() {
	var appendArray []int

	start := time.Now()
	for i := 0; i < 100000; i++ {
		appendArray = append(appendArray, i)
	}
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("Time elapsed for appending: ", elapsed)

	var prependArray []int

	start2 := time.Now()
	for i := 0; i < 100000; i++ {
		prependArray = append([]int{i}, prependArray...)
	}
	elapsed2 := time.Since(start2).Milliseconds()
	fmt.Println("Time elapsed for prepending: ", elapsed2)

} */
