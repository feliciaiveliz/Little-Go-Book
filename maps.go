package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["Felicia"] = 28
	occupation, exists := lookup["engineer"]

	fmt.Println(occupation, exists) // 0 false
	fmt.Println(len(lookup))        // 1
	delete(lookup, "Felicia")       // map[]
	fmt.Println(lookup)             // []

	person := make(map[string]int, 10)
	fmt.Println(len(person))
}

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan // string key Saiyan value
}

func main() {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
}

// iterate over map

func main() {
	lookup := map[string]int{
		"felicia": 28,
		"vinton":  29,
	}

	for key, value := range lookup {
		fmt.Println(key, value)
	}
}
