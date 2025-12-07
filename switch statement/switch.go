package main

import (
	"fmt" // Package for formatted I/O
	"time"
)

// const (
// 	i := 1 // const kai andar shorthand declaration nhi hota hai
// )
func main() {
	// switch statement
	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	default :
		fmt.Println("other")
	
	}
	mult()
}

// Multiple case values in a single case
func mult() {
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
			fmt.Println("its Weekend")
	default:
			fmt.Println("Its weekday")
	}
}