package main

import "fmt"

func main() {
	func() {
		fmt.Println("Why do this?")
	}()
}
