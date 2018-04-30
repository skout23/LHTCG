package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good Mythincal Morning",
		1: "Howdy do?",
		2: "Yo! Yo!",
		3: "Sup?",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
