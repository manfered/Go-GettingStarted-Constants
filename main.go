package main

import "fmt"

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
}
