package main

import "fmt"

func main() {
	o := [5]int{1, 2, 3, 4, 5}
	i := [...]int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	 fmt.Println("Array:", o)
	 fmt.Println("Array with implicit length:", i)
}