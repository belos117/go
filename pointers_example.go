package main

import "fmt"

// Person struct represents a person with a name and age
type Employee struct {
	Name string
	Age  int
}

func UpdateAge(a *Employee, newAge int){
	a.Age = newAge
}

func main() {
	// Create Employee instances
	emp := Employee{Name: "Henkel", Age: 30}

	UpdateAge(&emp, 35) // Pass the address

	fmt.Println("Your new age is:", emp.Age) // Access the updated age
}
