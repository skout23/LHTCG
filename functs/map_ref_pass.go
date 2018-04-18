package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Scott"]) // 43

}

func changeMe(z map[string]int) {
	z["Scott"] = 43
}
