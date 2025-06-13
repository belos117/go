package main

import "fmt"

func main() {
	s := 10
	var t *int = &s // Pointer to s
	fmt.Println("Hello, Go World", *t) // Dereferencing the pointer to get the value of s
	fmt.Println("Value of s:", s)
	fmt.Println("Address of s:", &s) // Address of s
	fmt.Println("Address stored in t:", t) // Address stored in pointer t
}