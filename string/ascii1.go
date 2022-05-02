package main

import "fmt"

func main() {
	lowercase := "abcdefghijklmnopqrstunwxyz"
	for _, c := range lowercase {
		fmt.Print(c, " ")
	}

	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range uppercase {
		fmt.Print(c, " ")
	}

	numbers := "0123456789"
	for _, n := range numbers {
		fmt.Print(n, " ")
	}

	fmt.Println("..............charcter of ascii values...........")
	sampleASCIIDigits := []int{97, 98, 99}
	for _, digit := range sampleASCIIDigits {
		fmt.Printf("Char %s\n", string(digit))
	}
}
