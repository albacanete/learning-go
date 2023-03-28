package main

import "fmt"

func main() {
	x, y := 5, 6
	a := add(x, y)
	fmt.Println(a)

	// call recursive function
	fmt.Println(factorial(5))

	// you can add an unlimited number of arguments to a function
	fmt.Println(addemup(10, 20, 30, 40, 50))
}

func add(num1, num2 int) int {
	return num1 + num2
}

// recursive function
func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
	/* it can also be
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
	*/
}

func addemup(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}
