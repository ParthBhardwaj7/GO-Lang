package main
import (
	"fmt" // Package for formatted I/O
)
// go lang mai ternary operator nhi hota hai if else se hi kaam chalana padta hai
func main() {
	// if else statement 
	num := 10
	if num >= 10 {
		fmt.Println ("Number is greater than or equal to 10")
	}else {
		fmt.Println("Number is less than 10")
	}

	// if else if else statement
	age :=10
	if age<18{
		fmt.Println("Minor") // agr age 18 se chota hai to Minor print hoga aur code aaga execute he nhi hoga
	}else if age>=18 && age<60{// || means OR && means AND
		fmt.Println("Adult")
	}else{
		fmt.Println("Senior Citizen")
	}

	//logical operators in if else
	var role = "admin"
	var hasPermission = true
	if role == "admin" || hasPermission {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}

	//if with a short statement
	if age:= 15; age < 18 { //age yaha sirf if ke andar hi valid hai
		fmt.Println("Teenager")
	} else {
		fmt.Println("Not a Teenager")
	}

	//Range OVer a slice 
	numbers := []int{10,20,30,40,50}
	for index, value := range numbers {
		fmt.Println(index,value)
	}

	//Ignoring index
	for _, value := range numbers {
		fmt.Println(value)
	}

	//Ignoring value
	for index := range numbers {
		fmt.Println(index)
	}
}

