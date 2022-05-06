package main

import "fmt"

type employee struct {
	name    string
	age     int
	salary  int
	address address
}

type address struct {
	state   string
	country string
}

func main() {
	address := address{state: "TamilNadu", country: "India"}
	emp := employee{name: "Sabari", age: 22, salary: 20000, address: address}
	fmt.Println(emp)
	fmt.Printf("state: %s\n", emp.address.state)
	fmt.Printf("Country: %s\n", emp.address.country)
	fmt.Println("name :", emp.name)
	//dot operator
	emp.name = "Arun"
	fmt.Println("name :", emp.name)

}
