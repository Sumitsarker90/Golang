package main

import "fmt"



func main(){

	// // Variable declearation in go 
	// var age int
	// var cgpa float32
	// var country string
	
	// // variable initialization
	
	
	// age =24
	// cgpa=3
	// country="Bangladesh"


// 	// Directly variable initialize

// 	var studentname="Sumit"
// 	// variable declearation without var
// 	studentAge :=32
//    const UNIVERSITY= "Aiub"
//    ///

   /// Input: scan, scanln, scanf


//    var fullName string
// //    var age int

// fmt.Print("Enter Your Name: ")
//    fmt.Scan(&fullName)
   

//  var num1,num2 int
//  fmt.Print("Enter Two numbers: ")
//  fmt.Scan( &num1, &num2)
//  fmt.Println(num1+num2)
//  fmt.Printf(" %v is the student name",fullName)

   
	// fmt.Printf("%v is a student and age is %v  and university is %v\n",studentname,studentAge,UNIVERSITY)
	// fmt.Println(age,cgpa, country,studentname,studentAge)

	var base,height,area float32

	fmt.Print("Enter Base= ")
	fmt.Scan(&base)

	fmt.Print("Enter Height= ")
	fmt.Scan(&height)
	
	
	area=0.5*base*height;

	fmt.Print("The area is: ",area)

	test("welcome sumit")

	func test()  {
      fmt.Print()
	}


}






