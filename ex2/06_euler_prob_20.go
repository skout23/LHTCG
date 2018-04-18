/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import "fmt"

func factorial(x int) int {

	if x == 1 {
		return 1
	}
	// fmt.Println("checking: ", x)
	return x * factorial(x-1)
}

//func sliceSum(n int) int {
//
//}

func main() {
	x := factorial(100)
	fmt.Println("factorial: ", x)
	// sum := sliceSum(x)

}
