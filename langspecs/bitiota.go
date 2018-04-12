package main

import "fmt"

const (
	_ = iota // 0
	// KB exports size of KB
	KB = 1 << (iota * 10) // 1 << (1 * 10))
	// MB exports size of MB
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	// GB export size of GB
	GB = 1 << (iota * 10) // 1 << (3 * 10)

)

func main() {
	fmt.Println("Binary\t\t\t\tDecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
