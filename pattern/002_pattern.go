package main

import "fmt"

func main() {
	var i, j, row, col int
	var a1, b0 string
	fmt.Println("Enter the row and col values")
	fmt.Scanln(&row, &col)
	fmt.Println("Enter the inside and outside values")
	fmt.Scanln(&a1, &b0)
	for i = 1; i <= row; i++ {
		for j = 1; j <= col; j++ {
			if i%2 == 0 {
				fmt.Printf("%s", a1)
			} else {
				fmt.Printf("%s", b0)
			}
		}
		fmt.Println()
	}
}
