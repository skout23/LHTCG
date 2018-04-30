package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good Mythical Moring",
		"Jenny": "Time for your spanking",
	}

	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting)

	myGreeting["Harleen"] = "Damn Son!"
	fmt.Println(myGreeting)

}
