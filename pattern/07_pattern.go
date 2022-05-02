package main

import "fmt"

func main() {
	var i, j, row, col int
	fmt.Println("enter the row and colom")
	fmt.Scanln(&row, &col)
	fmt.Println(...........Rectangle.................)
	for i = 1; i <= row; i++ {
		for j = 1; j <= col; j++ {
			fmt.Print("l")
		}
		fmt.Println()
	}

}
