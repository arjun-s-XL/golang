package main

import "fmt"

func split(input int) (x, y int) {
	x = input * 10
	y = input * 5
	return
}

func main() {
	var input int
	fmt.Println("Enter input: ")
	fmt.Scan(&input)

	fmt.Println("The modified values are")
	fmt.Print(split(input))
}
