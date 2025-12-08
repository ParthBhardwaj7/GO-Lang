package main

import "fmt"

// "fmt"

// func add(a int, b int)int {// isko add(a,b int) bhi likh sakte hai
// 	return a + b

// }

// func getLanguages()(string,string,string){
// 	return "golang","javascript","c"
// }


// func processIt(fn func(a int)int)int{
// 	return fn(2)
// }
func processIt() func(a int)int {
	return func(a int)int {
		return a*a
	}
}
func main() {
	// fn := func(a int)int {
	// 	return a*a
	// }
	fn:=processIt()
	result:=fn(6)
	fmt.Println("Result:",result) 
	// result := add(5, 10)
	// fmt.Println("Sum:", result)
	// lang1,lang2,lang3 := getLanguages()
	// fmt.Println(lang1,lang2,lang3)
}	