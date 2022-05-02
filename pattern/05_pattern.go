package main

import "fmt"

func main() {
	var i, j, row int
	fmt.Println("enteer the row")
	fmt.Scanln(&row)
	fmt.Println("........K shap Patterrn..............")
	for i = row; i >= 1; i-- {
		for j = 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	for i = 2; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println(".............Right Arrow Number Pattern..........")
	for i = 1; i <= row; i++ {
		for j = 1; j <= row; j++ {
			if j < i {
				fmt.Print("")
			} else {
				fmt.Print(j)
			}

		}
		fmt.Println()
	}
	for i = 1; i < row; i++ {
		for j = 1; j <= row; j++ {
			if j < row-i {
				fmt.Print("")
			} else {
				fmt.Print(j)
			}

		}
		fmt.Println()
	}
}
