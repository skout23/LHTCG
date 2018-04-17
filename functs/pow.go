package main

import "fmt"
import "math"

func power(x int) func(int) int {
	return func(y int) int {
		return int(math.Pow(float64(x), float64(y)))
	}
}

func main() {
	var pow2 = power(2)
	fmt.Printf("2 ^ 5 = %d\n", pow2(5))
	var pow3 = power(3)
	fmt.Printf("3 ^4 = %d\n", pow3(4))
}
