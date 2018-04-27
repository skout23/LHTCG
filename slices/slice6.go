package main

import "fmt"

func main() {

	customerNumber := make([]int, 3)
	//  3 is length & cap ?
	// // length - number of elements referred to by slice
	// //  capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by slice
	// 5 is capacity - number of elements in the underlying attay
	// you could also do it like this.

	greeting[0] = "Hello"
	greeting[1] = "Konichiwa"
	greeting[2] = "Guten Nicht"
	// panic: runtime error: index out of range
	// greeting[3] = "Guten BAD"

	fmt.Println(greeting[0])
	fmt.Println(greeting[1])
	fmt.Println(greeting[2])
	// fmt.Println(greeting[3])

	greeting = append(greeting, "Guten BAD")
	fmt.Println(greeting[3])

}
