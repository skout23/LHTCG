package main

import "fmt"

const p string = "Death & Taxes"

func main() {
	const q int = 42

	// shady // p := "something else"

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}
