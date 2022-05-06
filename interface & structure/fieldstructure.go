package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}
type emplo struct {
	BOY int
	DOB int
}

func main() {
	emp := employee{Name: "Sabari", Age: 30, Salary: 2000}
	emp1 := employee{"Arun", 25, 2000}
	//Converting to jsonn
	empJSON, err := json.MarshalIndent(emp, " ", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))

	empJSON1, err := json.MarshalIndent(emp1, " ", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON1))
	emp2 := emplo{BOY: 1998, DOB: 27}
	fmt.Println(emp2)

}
