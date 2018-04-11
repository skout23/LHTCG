package main

import "fmt"

func main() {
	for index := 0; index < 200; index++ {
		fmt.Printf("%d \t %b \t %x \t %q \t\n", index, index, index, index)
	}
}
