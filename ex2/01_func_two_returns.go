package main

import "fmt"

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	n, t := half(43)
	if t {
		fmt.Println(n, "even")
	} else {
		fmt.Println(n, "not even")
	}
}
