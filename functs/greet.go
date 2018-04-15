package main

import "fmt"

func main() {
	greet("Scott")
	greet("Paul")

}

func greet(name string) {
	fmt.Println(name)
}
