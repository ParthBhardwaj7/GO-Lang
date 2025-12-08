package main

import (
	"fmt"
	"maps"
)

func main() {
	//Map is collection of key value pairs
	//Map is reference type
	// maps -> hash table, dictionary, object in other languages
	m := make(map[string]int) //

	//setting an element in map
	m["name"] = 25
	m["area"] = 55
	//get an element
	fmt.Println(m["name"],m["area"])
	fmt.Println(m["age"]) // if key not present returns zero value of value type

	ms := map[string]int{
		"john": 30,
		"doe":  40,
	}
	//delete(ms,"doe")// delete key value pair
	//clear(ms) // clear all key value pairs

	fmt.Println(ms)
	mq:= map[string]int{"alice": 28,"bob":   34,}
	fmt.Println(mq)

	v ,ok := mq["bob"] // check if key present By ok Joh ki value true hoga aur agar key present nahi hoga to false hoga
	// v means value associated with key
	// _ means blank identifier mtlb mujhe value ki jarurat nahi hai
	fmt.Println(v)
	if ok {
		fmt.Println("key present")
	} else {
		fmt.Println("key not present")
	}

	m1 := map[string]int{"x": 10, "y": 20}
	m2 := map[string]int{"x": 10, "y": 20}

	//fmt.Println(m1 == m2) // maps cannot be compared directly except for nil

	fmt.Println(maps.Equal(m1, m2)) // using maps.Equal function to compare two maps
}	