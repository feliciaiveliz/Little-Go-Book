// go doesn't have objects, inheritance, polymorphisms, overloading
// structs are similar to classes
// outlines for things
package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

// func main() {
// 	goku := Saiyan{"Goku", 9000} // create goku struct with name "Goku", power 9000
// 	Super(goku)                  // pass to Super, which increases power to 19000 -- of the copy of 'goku'
// 	fmt.Println(goku.Power)
// }

// func Super(s Saiyan) {
// 	s.Power += 10000
// }

// passing a pointer

func main() {
	goku := &Saiyan{"Goku", 9000} // get address of value -- address of operator
	Super(goku)
	fmt.Println(goku.Power) // 19000
}

func Super(s *Saiyan) { // pass the address of type *Saiyan, so *Saiyan is pointer to value of type
	s.Power += 10000 // now we can change the value held at that address
}

// think of it as copy of address or pointer of the actual value
