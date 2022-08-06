package main

import "fmt"

func zero(x *int) { // what it expects (memory address)
	*x = 0 // change the value of 'x' to 0
}

func main() {

	x := 5
	fmt.Println(x) // 5
	zero(&x)
	fmt.Println(x) // 0
}

package main 

import "fmt"

type Person struct {
  name string
	age int
}

func (p *Person) changeName(name string) {
  p.name = name
}

func (a *Person) addYear() {
  a.age++ 
}

func (p *Person) getAge() int {
	return p.age
}


func main() {
  kim := Person{"Kim", 34}
	kim.changeName("Felicia")
	kim.addYear()
	fmt.Println(kim.age)
}