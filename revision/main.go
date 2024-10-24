package main

import (
	"crypto/rand"
	"fmt"
	"time"
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
getRandom()

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

func getRandom() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day: ", dow)
	var result string
	switch dow {
	case 1:
	result = "It's Sunday"
	case 2:
	result = "It's Monday"
	default:
	result = "It's some other day"
	}
	fmt.Println(result)
   }
	
