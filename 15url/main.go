package main

import (
	"fmt"
	"net/url"
)

const myurl string ="https://lco.dev:3000/learn?coursename=reactjs"

func main() {

	fmt.Println("Welcome to golang url")
	fmt.Println(myurl)

	// parsing

	result, _ :=url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)


}
