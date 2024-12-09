package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = v.X * 2
	p.Y = v.Y * 4
	fmt.Println(p.X, p.Y)

}
