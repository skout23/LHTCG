package main

import "fmt"

func main() {

	switch "Medhi" {
	case "Daniel":
		fmt.Println("Sup Dan")
	case "Medhi":
		fmt.Println("My Dog Medhi")
	case "Jenny":
		fmt.Println("Hello Dear")
	default:
		fmt.Println("Who goes there?")
	}
}
