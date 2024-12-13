package main

import "fmt"

func main() {
	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	q := p[:2]
	r := p[3:5]
	s := p[4:]
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
}
