package main

import (
	"fmt"
)

//comment
/*single line*/
func man(){
  fmt.Println("GO is fun to learn!")/*display output*/

  //explicit declaration of variables
var aStringVariable string="I am Ian Njoroge"
fmt.Println(aStringVariable)
var agoodStringVariable string="This is another string variable"
fmt.Println(agoodStringVariable)

var names string="KPC workers"
fmt.Println(names)
var numb int=23
fmt.Println(numb)
//implicit declaration
 age :=67
fmt.Println(age)

var name string
var miaka int

fmt.Print("Enter Your name: ")
 fmt.Scanln(&name)
 fmt.Println("Hello ", name)
 fmt.Print("Enter Your age: ")
 fmt.Scan(&miaka)
 fmt.Println("Your age is ", miaka)
 
}