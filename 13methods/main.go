package main

import "fmt"


func main() {
	fmt.Println("Struct in golang")

	// no inheritance in glang; no super, no parent
    sumit := User{"sumit","sumit@gmail.com", true,16}
	fmt.Println(sumit)
	fmt.Printf("Sumit details are: %+v\n", sumit)
	fmt.Printf("Name is %v and email is %v: \n", sumit.Name,sumit.Email)

	 sumit.Getstatus()
	 sumit.Newmail()

	 fmt.Printf("Name is %v and email is %v: \n", sumit.Name,sumit.Email)

}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func(u User)Getstatus(){

	fmt.Println("Is user active: ",u.Status)
}


func (u User) Newmail(){
	u.Email="test@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}
