package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(67, 34, 56, 2, 78, 97, 7, 2, 1)
	fmt.Println(greatest)

}
