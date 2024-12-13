package main

import (
	"fmt"
	"math"
)

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	a := circleArea(10.0)
	fmt.Println(a)
}
