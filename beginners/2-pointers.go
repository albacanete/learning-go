package main

import "fmt"

func main() {
	x := 5

	fmt.Println(x)
	fmt.Println(&x)

	changeValue(x) // when called like this the value of x is not changed since it does not have scope outside of the main function
	fmt.Println("No pointer:", x)

	changeValuePointer(&x)
	fmt.Println("Pointer:", x)
}

func changeValue(x int) {
	x = 7
}

func changeValuePointer(x *int) {
	*x = 7
}
