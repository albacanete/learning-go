package main

import "fmt"

func main() {
	var name string = "hola que tal"

	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(name + " como va")

	// printing a float with precision 0.3
	const pi float64 = 3.1416
	fmt.Printf("%.3f \n", pi)

	// printing a type
	fmt.Printf("%T \n", name)

	// printing a boolean
	win := true
	fmt.Printf("%t \n", win)

	// printing an integer
	x := 5
	fmt.Printf("%d \n", x)

	// printing the binary of a number
	fmt.Printf("%b \n", 6)

	// printing characters associated to an ASCII code
	fmt.Printf("%c \n", 34)

	// printing hex code of a value
	fmt.Printf("%x \n", 15)

	// printing scientific notation
	fmt.Printf("%e \n", pi)
}
