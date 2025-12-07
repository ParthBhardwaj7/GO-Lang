package main

import (
	"fmt"
)

func main() {
	//Numbered Sequence of Specific Length
	var nums [4]int
	fmt.Println(len(nums)) // Array ki length

	nums[0] = 10
	fmt.Println(nums[0])
	fmt.Println(nums) // Arrays mai sirf ek he value di thi voh bhe 1st place ki toh uska alva sab zero ho jayegi 
	

	var vals [4]bool
	fmt.Println(vals)//Bollean kai case mai False hota hai
	vals[2]=true
	fmt.Println(vals)

	//string ka case mai 
	var name [3]string 
	name[0] ="go lang"
	fmt.Println(name) //type kai hissab sai zeroed value change hoti hai 
	
	//To declare in single line
	nums1 := [3]int{1,2,3}
	fmt.Println(nums1)

	//2d Arrays
	nums2 :=[2][2]int{{2,3}, {5,6}}
	fmt.Println(nums2)
	
	//Advance mai pata hai toh Use -Array
 	// memory optimize rhti hai 
        //constant time access
}