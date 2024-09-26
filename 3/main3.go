package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	name := [4]string{"yoshi", "mario", "peach", "browser"}
	name[1] = "ravin"
	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))

	// slices (use arrays under the hood)
	var scores = []int{100, 34, 243}
	scores[2] = 23

	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := name[1:3]
	fmt.Println(rangeOne)

	rangeTwo := name[2:]
	fmt.Println(rangeTwo)

	rangeOne = append(rangeOne, "koops")
	fmt.Println(rangeOne)

}
