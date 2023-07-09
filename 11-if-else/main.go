package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("if else in golang")

	loginCount := 12
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "something else"
	}

	fmt.Println(result)

	//if err !=nil{}


}
