package main

import "fmt"

func main() {
	age := 27
	name := "ravin"

	//print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("newline \n")

	fmt.Println("hello ravin!!")
	fmt.Println("goodbye")

	fmt.Println("my age is ", age, "and my name is", name)

	//Printf(formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you score %f points! \n", 22.55)
	fmt.Printf("you score %0.1f points! \n", 22.55)

	// Sprintf (save formatted strings)

	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}
