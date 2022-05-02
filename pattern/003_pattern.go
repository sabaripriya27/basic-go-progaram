package main

import "fmt"

func main() {
	var i, j, n, k int
	fmt.Println("enter n vlaues")
	fmt.Scanln(&n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
	for i = n - 1; i > 0; i-- {
		for j = 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
}
