package main
import (
	"fmt" // Package for formatted I/O
)
// Constants or Variable function kai bahar declare kiye jaa skta hain But shorthand declaration (:=) sirf function ke andar hi hota hai
// Const declarations main function ke bahar hote hain
const (
	PI       = 3.14               // Constant for Pi
	AppName  = "GoLangApp"        // Constant for Application Name
	MaxUsers = 100                 // Constant for Maximum Users
)
func main() {
	//local constants
	const LocalConst = "I am local to main function"
	fmt.Println(LocalConst)  //INSIDE const ≠ OUTSIDE const
	MaxUsers := 50          // Shadowing the global constant
	fmt.Println("Shadowed MaxUsers:", MaxUsers)
	fmt.Println("Constants:")
	fmt.Println("PI:", PI)
	fmt.Println("Application Name:", AppName)
	fmt.Println("Maximum Users:", MaxUsers)
}