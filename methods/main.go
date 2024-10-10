package main

import (
	"fmt"
)

func main(){
fmt.Println("Methods in golang")

Ianoh := User{"Ianoh","Ianoh.dev",true,19}
fmt.Println(Ianoh)
fmt.Printf("Ianoh's details are: %+v\n",Ianoh)
fmt.Printf("Name is %v.",Ianoh.Name)
Ianoh.NewMail()//method

}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
func ( u User) NewMail() {
	u.Email="njorogeian@gmail.com"
	fmt.Println("Email of new user is: ",u.Email)
}





