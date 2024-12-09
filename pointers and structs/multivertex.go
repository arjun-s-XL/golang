package main

import "fmt"

type vertex struct {
	X int
	Y int
}

var (
	v1 = vertex{1, 2}
	v2 = vertex{X: 2}
	v3 = vertex{Y: 3}
	p  = &v1
)

func main() {
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p.X)
}
