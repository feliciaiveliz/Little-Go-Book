// functions can return multiple values

func log(message string) { // no return value

}

func add(a int, b int) int { // one return value of int

}

func power(name string) (int, bool) { // two return values of int and bool
	value, exists := power("Felicia")

	if exists == false {
		// handle error
	}

	// or if you don't need one of the return values
	_, exists := power("Felicia")

	if exists == false {
		// handle error
	}
}

func add(a, b int) int { // shorter syntax for multiple parameters

}

  