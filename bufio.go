package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){

	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Hello, how are you?")
	input,_:=reader.ReadString('\n')
	fmt.Println("You've entered :",input)

	i1,i2,i3:=23,45,89
	intSum :=i1+i2+i3
	fmt.Print(intSum)

	

/*declaring pointers*/
value1 := 42
 var pointer1 = &value1/*using ampersand to assign*/
 fmt.Println("Value of pointer1: ", *pointer1)

 value3:=90
 var pointer2=&value3
 fmt.Print("Value of ponter 2",*pointer2)


 /*array declaration*/

 var array1 = []float32{10.5, 5.2, 2.88}
 var array2 [10]int
 var i, j int


 //Initializing elements of the array
 for i =0; i < 10; i++{
 array2[i] = i + 50 //setting element at location 
i to i+50
 }
 //Print the value of each elements of array1
 fmt.Println("Elements stored in Array1")
 for j =0; j <3; j++{
 fmt.Printf("Element[%d] = %f \n", j, array1[j])
 }
 //Print the value of each elements of array2
 fmt.Println("Elements stored in Array2)
 for j =0; j <10; j++{
 fmt.Printf("Element[%d] = %d \n", j, array2[j])
 }
 fmt.Println("Size of array1: ", len(array1))
 fmt.Println("Size of array2: ", len(array2))
}



}