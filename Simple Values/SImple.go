package main

import (
	"fmt"// Package for formatted I/O
)

func main() {
	//simple values
	//integers
	var a int = 10
	var b int64 = 20
	c := 30 //type inference or Short declaration
	fmt.Println("Integers:", a, b, c)

	//floating point
	var d float32 = 10.5
	var e float64 = 20.99
	f := 30.75
	fmt.Println("Floating Points:", d, e, f)	

	//boolean
	var h = true
	var g bool = true
	z := false
	fmt.Println("Booleans:", g, h, z)

	//string
	var Str3 = "Welcome to GoLang"
	var str1 string = "Hello"
	str2 := "GoLang"
	fmt.Println("Strings:", str1, str2, Str3)

	//complex numbers
	var complex1 complex64 = 1 + 2i
	var complex2 complex128 = 2 + 3i
	complex3 := 3 + 4i
	fmt.Println("Complex Numbers:", complex1, complex2, complex3)

	//rune and byte
	var r rune = 'a' //Rune = ek Unicode character// Means: world ke saare characters (Hindi, emoji, Arabic, etc.) support karta hai
	var bte byte = 65 //Byte = 0 to 255 tak ke numbers ko represent karta hai
	r2 := 'b' 
	fmt.Println("Rune and Byte:", r, bte, r2)
}