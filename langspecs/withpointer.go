package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) // address in func zero
	fmt.Println(&x)        // addres in func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in func main
	fmt.Println(&x)        // addres in func main
	zero(x)
	fmt.Println(x) // still 5
}
