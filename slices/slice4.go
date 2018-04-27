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

	for i, currentEntry := range greetings {
		fmt.Println(i, currentEntry)
	}
	for j := 0; j < len(greetings); j++ {
		fmt.Println(greetings[j])
	}
}
