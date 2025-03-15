package main
import "math/rand"
func randomInt(min int, max int) int {
    return rand.Intn(max - min + 1) + min
}
func main() {
    fmt.Println("The random number is", randomInt(1, 10))
}
