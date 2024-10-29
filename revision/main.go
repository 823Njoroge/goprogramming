package main

import (
	"fmt"
)

// Function to calculate and print the perimeter of a rectangle
func getPerimeter(width, length int) {
	perimeter := 2*(width + length)
	fmt.Println("The perimeter is:", perimeter)
}

/*/ Main function
func main() {
	}
		width := 24
	length := 23

	// Calculate and print the area
	area := length * width
	fmt.Println("The area of the rectangle is:", area)

	// Call other functions
	getPerimeter(width, length)
	checkPrime()
	getRandom()
}

// Function to check if a number is prime
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

// Function to print prime numbers from 1 to 100
func checkPrime() {
	fmt.Println("Prime numbers between 1 and 100:")
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// Function to generate and print a random day of the week
func getRandom() {
	rand.Seed(time.Now().UnixNano()) // Seed the random generator

	day := rand.Intn(7) + 1 // Random number between 1 and 7
	fmt.Println("Random day number:", day)

	var result string
	switch day {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	case 3:
		result = "It's Tuesday"
	case 4:
		result = "It's Wednesday"
	case 5:
		result = "It's Thursday"
	case 6:
		result = "It's Friday"
	case 7:
		result = "It's Saturday"
	default:
		result = "Invalid day"
	}

	fmt.Println(result)
}*/

//Handson challenge
// isOdd checks if a number is odd
func isOdd(num int) bool {
	return num%2 != 0
}

// filter applies the isOdd function to each number in the slice
func filter(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if isOdd(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	// Define the slice of numbers (fixed the syntax error in the original slice)
	numbers := []int{3, 2, 4, 5, 6, 7, 8, 9, 8}

	// Filter odd numbers
	oddNumbers := filter(numbers)

	// Print all numbers and whether they're odd or even
	fmt.Println("Checking all numbers:")
	for _, num := range numbers {
		if isOdd(num) {
			fmt.Printf("%d is an odd number\n", num)
		} else {
			fmt.Printf("%d is an even number\n", num)
		}
	}

	fmt.Println("\nFiltered odd numbers:", oddNumbers)
}
