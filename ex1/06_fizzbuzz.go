package main

import "fmt"

func main() {
	count := 100
	sum := 0
	for i := 0; i < count; i++ {
		if (i%3 == 0) && (i%5 != 0) {
			fmt.Println(i, ": Fizz")
			sum = sum + i
		}
		if (i%3 != 0) && (i%5 == 0) {
			fmt.Println(i, ": Buzz")
			sum = sum + i
		}
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println(i, "Fizzbuzz")
			sum = sum + i
		}
	}
	fmt.Println("And the total sum is:", sum)
}
