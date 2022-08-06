// arrays are fixed and cannot change size

package main

import "fmt"

func main() {
	scores := [4]int{9001, 9333, 212, 33}
	fmt.Println(len(scores))

	for _, value := range scores {
		fmt.Println(value)
	}
}

package main

import "fmt"

func main() {
  scores := []int{1, 2, 3, 4, 5}
	scores = scores[:len(scores) -1]
	fmt.Println(scores)
}

package main

import (
	"fmt"
	"math/rand"
	// "sort"
)

func main() {
	scores := make([]int, 100)
	
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	fmt.Println(scores)

	firstFive := make([]int, 5)
	copy(firstFive, scores[:5])
	fmt.Println(firstFive)
}


func main() {
	scores := make([]int, 100)
	
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	fmt.Println(scores)

	firstFive := make([]int, 5)
	copy(firstFive[2:4], scores[:5]) // only copies 3rd and 4th ints from 'scores' and leaves 0 for rest
	fmt.Println(firstFive)
}


func main() {
	scores := make([]int, 100)
	
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	fmt.Println(scores)

	firstFive := make([]int, 5)
	copy(firstFive, scores[:7]) // when past the length limit, will still only copy the exact length specified during creation
	fmt.Println(firstFive)
}