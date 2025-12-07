//Main package hai jisme application ka entry point hota hai
// controllers → package controllers  Application ke core logic ke liye Code hota hai
// utils → package utils  YEh helper Code hota hai 
// db → package db  Database se related Code hota hai
// handlers → package handlers  http request handlers ke liye Code hota hai
// models → package models Data structures aur business logic ke liye Code hota hai

package main

import (
	"fmt"// Package for formatted I/O
	"time"
)

func main() {
	fmt.Println("Server starting...")
	
	// Simulating a task
	time.Sleep(1 * time.Second)// Sleep for 1 second to simulate work
	
	fmt.Println("Hello, Parth! Your Go environment is ready.")
	fmt.Println("Current Time:", time.Now().Format(time.RFC1123))
}