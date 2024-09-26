package main

import "fmt"

func main() {
	fmt.Println("Hello Ravin !!")

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "browser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory

	var numOne int8 = 25
	var numTwo int8 = -124
	var numThree uint8 = 34

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 35.98
	var scoreTwo float64 = 34567865432.4356
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
