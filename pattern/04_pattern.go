package main

import "fmt"

func main() {
	var i, j, row, number int
	fmt.Println("enter the starting number")
	fmt.Scanln(&number, &row)
	for i = 1; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(number)
			number++
		}
		fmt.Println()
	}
}
