// project path C:\Users\Admin\Desktop\Go programming\codes
package main

import "fmt"
	

type employees struct{
	eName string
	eNumber int
	eLocation string
	
	
	}

func main(){


var Employee1 employees
var Employee2 employees

Employee1.eName="John mwangi"
Employee2.eNumber=345
Employee1.eLocation="Nairobi"

fmt.Printf("The name of the first employee is:%S\n",Employee1.eName)


theAnswer := -89
 var result string
 if theAnswer < 0 {
 result = "Less than Zero"
 } else if theAnswer == 0 {
 result = "Equal to Zero"
 } else {
 result = "Greater than Zero"
 }
 fmt.Println(result)


 total := 0
 for k := 0; k < 5; k++ {
 total += k
 }
 fmt.Println(total)

 if total==10{
	fmt.Printf("The total is 10")
 }else{
	goto theEnd
 }
 theEnd: fmt.Printf("End of program")
}







	

	