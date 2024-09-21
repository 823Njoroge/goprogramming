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
}