package main

import "fmt"

func main() {
	r := map[string]int{
		"Henkel": 30,
		"John":   25,
		"Jane":   28,
	}
	fmt.Println("Hello, Go World", r["Henkel"])
	fmt.Println("Henkel, How are you? You are", r["Henkel"], "years old.")
	fmt.Println("John, How are you? You are", r["John"], "years old.")
	fmt.Println("Jane, How are you? You are", r["Jane"], "years old.")
	fmt.Println("Total number of people:", len(r))
	fmt.Println("All people in the map:")
}