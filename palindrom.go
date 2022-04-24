package main

import "fmt"

func palindrom(num int) int {

	var temp, sum, start, end, remainder int

	fmt.Print("Enter the Minimum and Maximum limit of Palindrome = ")
	fmt.Scanln(&start, &end)

	fmt.Print("Palindrome Numbers between ", start, " and ", end, " are : ")
	for num = start; num <= end; num++ {
		sum = 0
		for temp = num; temp > 0; temp = temp / 10 {
			remainder = temp % 10
			sum = sum*10 + remainder
			if temp == 0 {
				break
			}
		}
		if num == sum {
			fmt.Print(num, "\t")
		}
	}
	return sum
}
func main() {
	var num int
	palindrom(num)
}
