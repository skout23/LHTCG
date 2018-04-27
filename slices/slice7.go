package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// length 	- 3
	// capcity 	- 5

	greeting[0] = "Greeting 0"
	greeting[1] = "Greeting 1"
	greeting[2] = "Greeting 2"

	greeting = append(greeting, "Greeting 4")
	greeting = append(greeting, "Greeting 5")
	greeting = append(greeting, "Greeting 6")
	greeting = append(greeting, "Greeting 7")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

}
