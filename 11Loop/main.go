package main

import "fmt"

func main() {

	days := []string{"Sunday", "Tuesday", "wednesday", "Friday"}
	fmt.Println("Length of Loop: ", len(days))
	fmt.Println("Length of Loop: ", days)

	for i:=0; i<len(days); i++{
			fmt.Println(days[i])
	}

	for i :=range days{
		fmt.Println(days[i])
	}

	for  index,day :=range days{
		fmt.Printf("index is %v and value is %v \n",index, day)
	}

	rougeValue :=1
	for rougeValue <10{
		if rougeValue == 5{
			break
		}
		fmt.Println("value is :",rougeValue)
		rougeValue++
	}

	

}
