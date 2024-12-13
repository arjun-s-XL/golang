package main

import "fmt"

func main() {
	var array [10]int
	for i := 1; i <= 10; i++ {
		array[i-1] = i
	}
	fmt.Println(array)

	primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	fmt.Println(primes)
}
