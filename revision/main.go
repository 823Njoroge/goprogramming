package main

import (
	"fmt"
)
func getPerimeter(width,length int){
	perimeter:=width*2 + length*2
	fmt.Print("The perimeter is:",perimeter)
	
}
//first challenge
func main() {
	width := 24
	length := 23
	
		area:= length *width
		fmt.Print("The area of the rectangle is",area)

	getPerimeter(width,length)
checkPrime()

}
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func checkPrime() {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
	
