package main

import "fmt"

func main() {
	// Example 1
	var name string
	name = "John"

	var height int
	height = 170

	fmt.Printf("Result Example 1:\nname: %s type: %T\nheight: %d type: %T\n", name, name, height, height)

	// Example 2
	var firstName string = "John"
	var lastName = "Doe"

	var favoritNumber float64 = 10.55
	var bodyWeight = 60.5

	fmt.Printf("Result Example 2:\nfirstName: %s\nlastName: %s\nfavoritNumber: %.2f\nbodyWeight: %.1f\n", firstName, lastName, favoritNumber, bodyWeight)

	// Example 3
	fullName := "John Doe"
	age := 27

	fmt.Printf("Result Example 3:\nfullName: %s\nage: %d", fullName, age)

	// Example 4
	var (
		a int     = 10
		b int     = 20
		c string  = "Hello"
		d bool    = true
		e float64 = 120.5
	)

	fmt.Printf("Result Example 4: a: %d b:%d c:%s d:%t 3:%.1f", a, b, c, d, e)
}
