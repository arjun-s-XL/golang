package main

import "fmt"

func main() {
	names := [4]string{
		"Adharshna",
		"Nithya",
		"Kesava",
		"Sanjpon",
	}

	proj := [4]string{
		"FMC",
		"UCSM",
		"INTS",
		"UCSM",
	}

	for i := 0; i < 4; i++ {
		fmt.Println(names[i], "works in", proj[i])
	}
}
