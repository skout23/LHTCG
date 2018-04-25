package main

import "fmt"

func main() {
	sup := (true && false) || (false && true) || !(false && false)
	fmt.Println(sup)
}
