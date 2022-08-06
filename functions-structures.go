package main

import "fmt"

type Saiyan struct {  // type from structs
	Name string
	Power int
}

func Super(*s) { // pass me the memory value of variable
	s.Power += 10000
}

// s -> memory address
// goku -> memory address  

goku := &Saiyan{"Goku", 9001} // goku -> memory address of struct {"Goku", 9001} 0xe97838430
// goku.Super()
Super(&goku)

goku := new(Saiyan)