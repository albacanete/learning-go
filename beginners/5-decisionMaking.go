package main

import "fmt"

func main() {
	age := 18
	fmt.Println("Age: ", age)

	// if-else statements
	if age > 18 {
		fmt.Println("Yes, you can vote!")
	} else { // else has to be on the line where if finishes
		fmt.Println("No, you can't vote!")
	}

	// switch statements
	age = 20
	fmt.Println("Age: ", age)

	switch age {
	case 16:
		fmt.Println("Prepare for college")
	case 18:
		fmt.Println("You can vote!")
	case 20:
		fmt.Println("Search for a job")
	default:
		fmt.Println("Still alive?")
	}
}
