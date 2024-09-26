package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i:", i)
	}

	names := []string{"modi", "mars", "ravin", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("Position at index %v is %v \n", index, value)
	}

	for index, value := range names {
		fmt.Printf("the position at index %v  is %v \n", index, value)
	}
	fmt.Print(names)
}
