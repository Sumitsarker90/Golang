package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")

	content := "This is need to go in a file- sumit file"

	file ,err :=os.Create("./myfile.txt")

	if err!=nil{
		panic(err)
	}

	length,err :=io.WriteString(file, content)
    if err!=nil{
		panic(err)
	}

	fmt.Println("Lenth is : ",length)
	defer file.Close()

	readFile("./myfile.txt")

}


func readFile(filname string){
	databyte,err := ioutil.ReadFile(filname)
	if err!=nil{
		panic(err)
	}

	fmt.Println("Text data insde the file is \n",string(databyte))

}

func checkNilerr(err error){
	if err!=nil{
		panic(err)
	}
}
