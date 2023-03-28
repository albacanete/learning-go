package main

import "fmt"

func main() {
	// Go only implements for loops
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	/* you can create you own while loop using for
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	*/

	// nested loop
	for i := 1; i < 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
