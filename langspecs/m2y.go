package main

import "fmt"

const (
	metersToYards float64 = 1.09361
	yardsToFeet   float64 = 3.00000
)

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	feet := yards * yardsToFeet
	fmt.Println(meters, "meters is ", yards, " yards, and ", feet, " feet.")
}
