package main

import (
	"fmt"
	"strings"
)

var walkOutside = true
var takeTheBluePill = false

func main() {
	fmt.Println("You find yourself in a dimlylit cavern. ")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	fmt.Println("There ist a sign near the entrance that reads 'No Minors!'.")
	var age = 26
	var minor = age < 18

	fmt.Printf()
}
