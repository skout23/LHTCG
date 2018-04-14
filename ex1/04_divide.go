package main

import "fmt"

func main() {

	var small_num int
	var large_num int
	fmt.Println("Give me a small number:")
	fmt.Scan(&small_num)
	fmt.Println("Give me a large number:")
	fmt.Scan(&large_num)
	fmt.Printf("The remainder is: %d\n", (large_num % small_num))
}
