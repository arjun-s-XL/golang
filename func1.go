package main

import "fmt"

func add(x int , y int) int{
	return x+y
}

func main() {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)
	
	fmt.Printf("The sum of %d and %d is %d",num1,num2,add(num1,num2))
}
