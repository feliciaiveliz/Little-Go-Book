package main

import "fmt"

func getPower() int {
	return 9001
}

// func main() {
// 	// var power int
// 	// power = 9000
// 	// var power int = 9000
// 	name, power := "Felicia", getPower() // declare variable and assign value to it
// 	fmt.Printf("%s's power is over %d\n", name, power)
// }

func main() {
	power := 1000 // declaring and assigning power to 1000
	fmt.Printf("default power is %d\n", power)

	name, power := "Felicia", 9000                     // declaring name and reassigning power
	fmt.Printf("%s's power is over %d\n", name, power) // power can only be an integer
}
