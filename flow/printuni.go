package main

import "fmt"

func main() {
	head := 5000
	tail := 6100
	for i := head; i < tail; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
