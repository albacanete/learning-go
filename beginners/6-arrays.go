package main

import "fmt"

func main() {
	var EvenNum [5]int

	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println("Manual:", EvenNum[2])

	// multiple base initialization
	Evens := [5]int{0, 2, 4, 6, 8}

	fmt.Println("Multiple:", Evens[2])

	// for loop --> _ means we are not using the iterator
	for _, value := range Evens {
		fmt.Println("Loop:", value)
	}

	// printing the for iterator
	for i, value := range Evens {
		fmt.Println("Iterator:", i, "Value:", value)
	}

	// slicing an array (you take a subpart)
	numSlice := []int{5, 4, 3, 2, 1}
	fmt.Println("Entire:", numSlice)
	sliced := numSlice[3:5] // 5 not included
	fmt.Println("Sliced:", sliced)

	// copying an array
	slice2 := make([]int, 5, 10)
	fmt.Println("slice2 before copy:", slice2)
	copy(slice2, numSlice)
	fmt.Println("slice2:", slice2)

	// appending to an array
	slice3 := append(numSlice, 3, 0, -1)
	fmt.Println("slice3:", slice3)
}
