package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "hello there friends"
	// strings package
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "Hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ll"))

	fmt.Println(strings.Split(greetings, " "))
	fmt.Println("original string value =", greetings)

	// sort package
	age := []int{43, 89, 34, 98, 23}

	sort.Ints(age)
	fmt.Println(age)

	index := sort.SearchInts(age, 89)
	fmt.Println(index)

	names := []string{"yogi", "modi", "ravin", "stupid"}

	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "ravin"))
}
