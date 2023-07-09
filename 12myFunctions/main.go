package main

import (
	"fmt"

)

func main() {
	greeter()

	result := adder(5,10)

	fmt.Println(result)

	proResult := proAdder(2,5,7,87)

	fmt.Println("Pro result is :", proResult)

	fmt.Println("welcome to functions in golang")
}


func greeter(){
	fmt.Println("Namastey golang")
}

func adder(valone int, valtwo int)int{
	return  valone + valtwo;
}

func proAdder(values...int) int {
	total := 0

	for _,val:= range values{
		total += val
	}

	return total
	
}