package main

import "fmt"

func main() {
	i, j := 42, 27

	p := &i
	fmt.Println(p)

	*p = j
	fmt.Println(*p)
	fmt.Println(&j)
}
