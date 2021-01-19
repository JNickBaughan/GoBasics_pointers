package main

import "fmt"

// the value is passed to this function so the original variable cannot be updated within the function
func zeroOutByValue(intValue int) {
	intValue = 0
	fmt.Println("from in the func ", intValue)
}

// a pointer is passed to this function so the original variable CAN be updated
func zeroOutByPointer(intPointer *int, updateValue int) {
	*intPointer = updateValue
	fmt.Println("from in the func ", *intPointer)
}

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

	zeroOutByValue(num);
	// the function cannot change 'num' here since the value is passed and not the actual variable
	fmt.Println(num)

	// when a function takes a pointer as a parameter it is able to affect the underlying variable
	// adding the '&' infront of a variable makes it a pointer to that variable
	zeroOutByPointer(&num, 3);
	fmt.Println(num)

	// or in this case pointerToNum already has a pointer to num
	zeroOutByPointer(pointerToNum, 8);
	fmt.Println(num)



}