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
	fmt.Println(myGreeting)

}
