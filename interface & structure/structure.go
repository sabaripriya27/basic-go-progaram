package main

import "fmt"

// Defining a struct type
type class struct {
	Name     string
	section  string
	standard string
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a class
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := class{"Akshay", "A", "3rd"}

	fmt.Println("Address1: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := class{Name: "Anikaa", section: "Ballia",
		standard: "3rd"}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := class{Name: "raj"}
	//by using dot operator
	a3.Name = "sabari"
	fmt.Println("Address3: ", a3)
}
