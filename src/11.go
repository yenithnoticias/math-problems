
// Generate a random number between 1 and 100
func RandomNumber(min, max int) int {
	return min + rand.Intn(max-min+1)
}