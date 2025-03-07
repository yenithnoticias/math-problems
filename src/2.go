 package main

import "fmt"

func generateRandomCode() {
	// Generate a random number between 1 and 10
	number := rand.Intn(10) + 1

	// Print the number to the console
	fmt.Println("Your random number is:", number)
}
