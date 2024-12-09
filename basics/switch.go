package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	hour := t.Hour()

	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Night")
	}
}
