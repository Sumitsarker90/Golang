package main

import (
	"fmt"

	
)

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"]="javascript"
	languages["rd"]="ruby"
	languages["py"]="python"

	fmt.Println("list of all languages",languages)
	fmt.Println("js shorts for: " , languages["js"])

	delete(languages,"rd")
	fmt.Println(languages)

	// Loops through maps

	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key,value)
	}




}
