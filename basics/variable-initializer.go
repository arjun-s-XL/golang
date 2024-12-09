package main

import "fmt"

var i, j int = 1, 2

func main() {
	var a, b int
	a = i
	b = j
	k := i + j
	fmt.Println(a, b, k)
}
