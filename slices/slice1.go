package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4]) // from [index:index -1]
	fmt.Println(mySlice[2])
	fmt.Println(string("myString"[2]))
}
