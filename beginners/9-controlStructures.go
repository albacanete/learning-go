package main

import "fmt"

func main() {
	FirstRun()
	SecondRun()
	fmt.Println()

	/* DEFER:
	- defers execution of a function until the surronding function returns
	- multiple defers are pushed into a stack and executed LIFO
	- generally used to cleanyp resources like a file or DB connection
	*/
	//fmt.Println("*** DEFER execution ***")
	//defer FirstRun()
	//SecondRun()

	fmt.Println(div(3, 0)) // should return panic since cannot divide by 0
	fmt.Println(3, 5)
	demPanic()
}

func FirstRun() {
	fmt.Println("I executed first")
}

func SecondRun() {
	fmt.Println("I executed second")
}

func div(num1, num2 int) int {
	defer func() {
		/* RECOVER:
		- helps regain normal flow of execution after panic
		- normally used with defer to recover panic in goroutine
		*/
		fmt.Println(recover())
	}()
	solution := num1 / num2

	return solution
}

/*
	PANIC:

- similiar to throwing a exception
- normally stops execution flow immediately
- defered functions are executed normally
*/
func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("Panic")
}
