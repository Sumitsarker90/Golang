package main

import "fmt"

const Pai float32= 3.1416 // public variable
	

func main()  {

	var unsername string = "Sumit"
	fmt.Println(unsername)
	fmt.Printf("Variable is of type: %T \n",unsername)

	var small uint8 = 255
		fmt.Println(small)
	fmt.Printf("Variable is of type: %T \n",small)

	var smallFloat float32 = 10.2323232323
		fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	var bigFloat float64 = 10.2323232323
		fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n",bigFloat)

	// default values and aliases 
	var var1 int
	fmt.Println(var1 )

	//  implicit type
	var website= "facebook.com"
	fmt.Println(website)
	
	// no var style
	numberofUser :=300
	fmt.Println(numberofUser)

	fmt.Println(Pai)
	fmt.Printf("Variable is of type: %T \n",Pai)

	
}