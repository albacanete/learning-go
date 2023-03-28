package main

import "fmt"

func main() {
	// used for hash maps
	StudentAge := make(map[string]int) // key: string, value: int

	StudentAge["Alba"] = 25
	StudentAge["Sara"] = 38
	StudentAge["John"] = 34
	StudentAge["Kate"] = 45
	StudentAge["Paul"] = 12
	StudentAge["Marc"] = 28
	StudentAge["Eli"] = 39

	fmt.Println("All students:", StudentAge)
	fmt.Println("Sara:", StudentAge["Sara"])
	fmt.Println("# of people:", len(StudentAge))

	// maps inside of maps
	superhero := map[string]map[string]string{
		"Superman": {
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman": {
			"RealName": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
