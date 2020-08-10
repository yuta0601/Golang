package main

import "fmt"

func main() {

	// AND
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", false && false)

	// OR
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || true =", false || true)
	fmt.Println("false || false =", false || false)

	// NOT
	fmt.Println("!true =", !true)
	fmt.Println("!false =", !false)
}
