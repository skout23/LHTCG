package main

import "fmt"

func main() {
	var x int
	fmt.Println("Give me a number", x)
	fmt.Scan(&x)
	fmt.Println("Checking...:", x)
	y := x % 2
	if y == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
