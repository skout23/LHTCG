package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Howdy World!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
