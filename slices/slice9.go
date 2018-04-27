package main

import "fmt"

func main() {
	mySlice := []string{"Mon", "Tues"}
	myOtherSlice := []string{"Wed", "Thur", "Fri"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)

}
