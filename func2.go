package main

import "fmt"


func swap(x string, y string)(string, string){
	return y,x
}


func main() {
	var input1, input2 string
	fmt.Println("Enter the first string: ")
	fmt.Scanln(&input1)

	fmt.Println("Enter the second string: ")
	fmt.Scanln(&input2)
	
	fmt.Println("The swapped values are: ")
	fmt.Println(swap(input1,input2))
}	