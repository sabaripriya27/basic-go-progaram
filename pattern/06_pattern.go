package main

import "fmt"

func main() {
	var i, j, k, row int
	fmt.Println("enter the row values")
	fmt.Scanln(&row)
	fmt.Println("..........Left Pascalss Number..........")
	for i = 1; i <= row; i++ {
		for j = i; j < row; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print(k)
		}
		fmt.Println()
	}
	for i = row; i >= 1; i-- {
		for j = i; j <= row; j++ {
			fmt.Print(" ")
		}
		for k = 1; k < i; k++ {
			fmt.Print(k)
		}
		fmt.Println()
	}
	fmt.Println("..........Right Pascalss Number..........")
	for i = 1; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(j)
		}

		fmt.Println()
	}
	for i = row - 1; i >= 1; i-- {
		for j = 1; j <= i; j++ {
			fmt.Print(j)
		}

		fmt.Println()
	}
}
