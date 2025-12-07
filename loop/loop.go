package main

import (
	"fmt" // Package for formatted I/O
)	
func main() {
	//for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}
	// In go lang for loop is used only for looping purpose
	//while loop (using for as while)
	j := 1
	for j <= 5 {
		fmt.Println("While Loop Iteration:", j)
		j++
	}
	//infinite loop (using for)
	k := 1 
	for {
		fmt.Println("Infinite Loop Iteration:", k)
		k++
	}
}