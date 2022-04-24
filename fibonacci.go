package main

import "fmt"

func printFibonacciSeries(num int) {
	fmt.Println("enter the n values")
	fmt.Scanf("%d", &num)
	a := 0
	b := 1
	c := b
	fmt.Printf("Series is: %d %d", a, b)
	for {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}

func main() {
	var n int

	printFibonacciSeries(n)

}
