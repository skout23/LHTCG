package main

import "fmt"

func main() {
	a := 44
	fmt.Println(a)  // 44
	fmt.Println(&a) // memory address of "a"

	var b *int = &a // pointer to an int, points at "a"

	fmt.Println(b)  // memory address of "a"
	fmt.Println(*b) // value stored in "a", the target of the pointer

}
