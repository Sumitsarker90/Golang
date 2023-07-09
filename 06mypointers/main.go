package main

import "fmt"

// variable are the reference of the memory

func main() {

	fmt.Println("Welcome to pointers")

	var ptr *int

	fmt.Println("THe value of pointers is: ", ptr)

	myNumber := 23

	var ptr1 = &myNumber                      // referencing  means (&)
	fmt.Println("value of pointer is", ptr1)  // pointer is the reference to the direct memory location
	fmt.Println("value of pointer is", *ptr1) // *pointer is refer the value of the memory address

	*ptr1 = *ptr1 + 2

	fmt.Println("The new value is : ", *ptr1)

}
