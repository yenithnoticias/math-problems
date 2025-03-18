
package main

import "math/rand"

func main() {
	// Generate a random number between 1 and 10
	n := rand.Intn(10) + 1

	// Print the result
	fmt.Println("The result is", n)
}