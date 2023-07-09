package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slice")

	// slice
	// How to create slice
	var fruitlist=[] string{"Apple","Tomato","peach"}
	fmt.Printf("Type of fruitlist is %T",fruitlist)
	fruitlist= append(fruitlist, "Mango","Banana")
	fmt.Println(fruitlist)

	fruit := append(fruitlist[1:3] )
	
	fmt.Println(fruit)

	highScores := make([]int, 4)
	highScores[0]=234
	highScores[1]=23
	highScores[2]=235
	highScores[3]=237

	fmt.Println(highScores)

	highScores=append(highScores, 555,666,222)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices based on index

	var courses =[]string{"recatjs", "javasript" ,"swift", "python"}
	fmt.Println(courses)

	index := 2
	courses=append(courses[ :index], courses[index+2:]...)
	fmt.Println(courses)
	

}
