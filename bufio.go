// project path C:\Users\Admin\Desktop\Go programming\codes
package main

import "fmt"

func main(){


	//slice declaration
	
	marks:=[]int{7,4,5,2,3,32,45,3}

	/*append() and copy()*/
	var nums []int
	nums=append(nums,10)
	fmt.Print("Nums=",nums)
	printItemsOfSlice(marks) //pass slice to a function

 if marks == nil {
 fmt.Printf("Slice is Nil")
 }
}
// function that accepts a slice and prints its details
 func printItemsOfSlice(x []int) {
 fmt.Printf("Length=%d Capacity=%d Slice=%v\n", 
len(x), cap(x), x)

//maps
new_map:=make(map[string]string)

new_map["Book1"]="Social studies"
new_map["Book2"]="Chem"
new_map["Book3"]="CRE"
new_map["Book4"]="English"
new_map["Book5"]="Math"

fmt.Println(new_map["Book1"])







}

	

	