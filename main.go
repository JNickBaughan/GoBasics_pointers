package main


import "fmt"

func main(){

	// this creates a type that is a pointer to an int that is named pointerToNum
	var pointerToNum *int

	num := 37
	pointerToNum = &num

	fmt.Println(num)
	// this displays as the memory address to the num variable
	fmt.Println(pointerToNum)
	// add the * to read num through the pointer and see the actual value
	fmt.Println(*pointerToNum)


}