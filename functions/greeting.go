package main

import (
	"fmt"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	//var name string
	//fmt.Println("Give me your name: ")
	//fmt.Scanln(&name)

	//sayGreeting(name)
	//sayBye(name)

	cycleNames([]string{"Stephen", "Jeremy", "Anthony"}, sayGreeting)
}
