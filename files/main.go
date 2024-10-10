package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content:="Add this to a file"

	file, err:=os.Create("./myfile.txt")

	if err !=nil{
		panic(err)
	}
length, err:=io.WriteString(file,content)
if err !=nil{
	panic(err)
}
fmt.Println("Length is:",length)
 defer file.Close()
 readFile("./myfile.txt")
} 



func readFile(filname string){
	databyte, err:=ioutil.ReadFile(filname)
	if err !=nil{
		panic(err)
	}
	fmt.Println("The data in the file is \n",string(databyte))
}
