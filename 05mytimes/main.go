package main

import (
	"fmt"
	"time"
)

func main()  {

	fmt.Println("Welcome to time study of golang")

	presentime := time.Now()
	fmt.Println(presentime)

	fmt.Println(presentime.Format("01-02-2006 15.04.05 Monday"))

	createtime := time.Date(1999,time.July,02,13,24,0,0,time.UTC)
	fmt.Println(createtime)
	fmt.Println(createtime.Format("01-02-2006 15.04.05 Monday"))

	
}
