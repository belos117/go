package main

import "fmt"

func main() {
	var pointArr [3]*int
	num1, num2, num3 := 1, 2, 3
	 pointArr[0] = &num1
	 pointArr[1] = &num2
	 pointArr[2] = &num3
	 for i := 0; i < len(pointArr); i++ {
	fmt.Println("Pointer value at index", i, ":", *pointArr[i])}
}
