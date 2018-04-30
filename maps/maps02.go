package main

import "fmt"

func main() {

	// declarative
	// var myGreeting = make(map[string]string)

	// declarative
	var myGreeting = map[string]string{}

	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

}
