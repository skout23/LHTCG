package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good monring!",
		1: "Bonjour!",
		2: "Que Tal?",
		3: "Sup?",
		4: "Howdy!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println(myGreeting)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)

	} else {
		fmt.Println("That value does not exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(myGreeting)

}
