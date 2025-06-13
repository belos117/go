package main

import "fmt"

// Employee struct
type Employee struct {
	Name   string
	Salary float64
}

// Update Salary function
func UpdateSalary(emp *Employee, newSalary float64) {
	emp.Salary = newSalary
}

// Update Salary function
func UpdateName(emp *Employee, newName string) {
	emp.Name = newName
}

func main() {
	// Create Employee instance
	emp := Employee{
		Name:   "Henkel",
		Salary: 50000.0,
	}

	// Update Salary
	UpdateSalary(&emp, 60000.0)

	// Update Name
	UpdateName(&emp, "Festus")

	// Print updated salary
	fmt.Println(emp)
	fmt.Printf("Updated Salary of %s: $%.2f\n", emp.Name, emp.Salary)
	fmt.Printf("Updated Name of Employee: %s\n", emp.Name)
}