package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------------")
	count := 1361
	for i := 0; i < count; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])

	}
}
