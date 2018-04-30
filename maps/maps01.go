package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k0"] = 7
	m["k1"] = 13

	fmt.Println("Map: ", m)

	v1 := m["k0"]
	fmt.Println("v1: ", v1)
	fmt.Println("Len: ", len(m))

	delete(m, "k1")
	fmt.Println("Map: ", m)

	_, prs := m["k1"]
	fmt.Println("Prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("MapN: ", n)

}
