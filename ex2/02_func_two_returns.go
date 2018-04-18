package main

import "fmt"

func main() {
	n := 43
	x, b := half(n)
	if b {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is not even")
	}

}

func half(z int) (int, bool) {
	var n int
	var t bool
	n = (z / 2)
	if n%2 == 0 {
		t = true
	}
	return n, t
}
