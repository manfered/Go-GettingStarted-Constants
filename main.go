package main

import "fmt"

// constant blocks
const (
	first  = 1
	second = "second"
)

// using iota at constant blocks
const (
	third  = iota
	fourth = iota
)

// iota will be resetted at new constant block
const (
	fifth = iota
	sixth
)

// constant expressions
const (
	seventh = iota + 7
	eighth
)

func main() {
	// implicit declaration of a constant
	const constantValueImplicit = 3
	// printing the constant value
	fmt.Println(constantValueImplicit)

	// printing the implicit constant value
	fmt.Println(constantValueImplicit + 2)
	fmt.Println(constantValueImplicit + 2.2)

	// explicit declaration of a constant
	const constantValueExplicit int = 3

	// printing the explicit constant value
	fmt.Println(constantValueExplicit + 2)
	// cast the constant value to float32
	// otherwise we would have compilation error
	fmt.Println(float32(constantValueExplicit) + 2.2)

	// using the constant values from constant block
	fmt.Println(first, second)

	// using iota
	fmt.Println(third, fourth)

	fmt.Println(fifth, sixth)

	fmt.Println(seventh, eighth)
}
