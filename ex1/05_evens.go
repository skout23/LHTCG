package main

import "fmt"

func main() {
	count := 100
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		}
	}
}
