package main

import "fmt"

func main () {
	type Person struct {
		Name string
		Age  int
}
var people[2]Person
 people[0] = Person{Name: "Henkel", Age: 30}
 people[1] = Person{Name: "Wilhelm", Age: 40}
 fmt.Println("People:", people)

 people2 := [...]Person{
  {Name: "Henkel", Age: 30},
  {Name: "Wilhelm", Age: 40}}
  fmt.Println("First Person:", people2[0].Age)
}