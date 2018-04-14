package main

import "fmt"

func main() {
	var myName string
	fmt.Println("Howdy, what's your name?")
	fmt.Scan(&myName)
	fmt.Println("Hello", myName)
}
