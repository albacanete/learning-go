package main

import "fmt"

func main() {
	// integers and floats
	var a int = 5
	var b float32 = 4.32
	fmt.Println(a, b)

	// constants
	const pi float64 = 3.1416
	fmt.Println(pi)

	// shorthand declaration
	x, y := 14, 15 // x = 14, y = 15
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("x + y = ", x+y)

	// multiple variables
	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB)

	// strings
	var Name string = "hola hello"
	fmt.Println(len(Name))

	// booleans
	// "long" version var isbool bool = true
	isbool := true
	hate := false
	fmt.Println("isbool:", isbool)
	fmt.Println("hate:", hate)
	fmt.Println("isbool && hate:", isbool && hate)
	fmt.Println("isbool || hate:", isbool || hate)
	fmt.Println("!isbool:", !isbool)
}
