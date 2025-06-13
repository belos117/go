package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Henkel", Age: 30}
	fmt.Println("Hello, Go World", p.Name)
	fmt.Println(p.Name, "How are you? You are", p.Age, "years old.")
}
// This code defines a struct called Person with fields Name and Age.
// It then creates an instance of Person, initializes it with values, and prints a greeting message using the struct's fields.