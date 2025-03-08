 package main
 
 import "fmt"
 import "math/rand"
 
 func randInt(min int, max int) int {
    return min + rand.Intn(max-min+1)
 }
 
 func solve(a, b float64) float64 {
    return math.Sqrt((a*a)+(b*b))
 }
 
 func main() {
    fmt.Println(solve(randInt(0, 10), randInt(0, 10)))
 }