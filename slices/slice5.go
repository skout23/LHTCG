package main

import "fmt"

func main() {

	greetings := []string{
		"Howdy",
		"This is nice",
		"I like this better than on 1 line",
		"I Wonder why we need a trailing comma",
		"on the last element of the slice",
	}

	fmt.Print("[1:2 ")
	fmt.Println(greetings[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greetings[:2])
	fmt.Print("[5:] ")
	fmt.Println(greetings[5:])
	fmt.Print("[:] ")
	fmt.Println(greetings[:])

}
