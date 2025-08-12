package main
import "fmt"

func main() {
	var numerator, denominator = 20, 2
	var result = intdiv(numerator, denominator)
	fmt.Println(result)
}
func intdiv(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}