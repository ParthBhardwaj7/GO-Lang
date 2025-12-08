package main

import (

	"fmt"
)

//iterating over data structures using range
//Loop Keyword hai jo iterate karta hai over array,slice,map,string etc
// Range returns two values index and value

func main() {
	nums := []int{6,7,8}
	for i :=0; i<len(nums); i++ {
		fmt.Println(nums[i])

	}

	// sum:=0

	// for _,num := range nums { // _ is blank identifier when we don't want index
	// 	sum=sum+num
	// 	fmt.Println(sum)
	// }

	for i,num := range nums {
		fmt.Println("index:",i,"value:",num)
	}

	m1:= map[string]int{"a":1,"b":2,"c":3}

	for k,v := range m1 {
		fmt.Println("key:",k,"value:",v)
	}
	for k := range m1 {
		fmt.Println("key:",k,)
	}

	//unicode code points rune
	//starting byte of rune 
	// 255 -> 1 byte , 2bytes -> 65535 , 3 bytes -> 16777215
	for i ,c := range "hello" {//agr i index chahiye to use kro warna _ use kro
	// i joh hai wo starting byte hai
	// i dekhega har character ko aur c us character ka rune value dega
	// i 0-255 -> 1 byte
	// i 256-65535 -> 2 bytes
	// i 65536-16777215 -> 3 bytes
	// yeh unicode kai hisaab se hota hai
		fmt.Println(i,c)
	}
	// agr humein sirf character chahiye to
	for _,c := range "hello" {
		fmt.Println(string(c))
	}

	





}