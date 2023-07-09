package main

import "fmt"

func main(){

	fmt.Println("Welcome to array in golangs")

	var fruitlist[4]string
	fruitlist[0]="apple"
	fruitlist[1]="apple"
	fruitlist[3]="apple"

	fmt.Println("Fruit list is",fruitlist)
	fmt.Println("Fruit list is",len(fruitlist))

	var colorlist=[3]string{"Black","white","green"}

	fmt.Println("Color list is",colorlist)
}
