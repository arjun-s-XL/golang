package main

import "fmt"

func main() {
	var input1 int
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&input1)

	var input2 float64 = float64(input1)
	fmt.Printf("%.2f", input2)

}
