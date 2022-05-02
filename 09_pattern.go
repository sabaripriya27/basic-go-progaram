package main

import "fmt"

func main() {
	var i, j, row, col int
	var a, b string
	fmt.Println("Enter the row and col values")
	fmt.Scanln(&row, &col)
	fmt.Println("Enter the inside and outside values")
	fmt.Scanln(&a, &b)
	for i = 1; i <= row; i++ {
		for j = 1; j <= col; j++ {
			if i == 1 || i == row || j == 1 || j == col {
				fmt.Printf("%s", a)
			} else {
				fmt.Printf("%s", b)
			}
		}
		fmt.Println()
	}
}
