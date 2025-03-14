import "math/rand"

func main() {
  rand.Seed(time.Now().UnixNano())
  x := rand.Intn(10)
  y := rand.Intn(10)
  fmt.Println("The sum of", x, "and", y, "is", x+y)
}