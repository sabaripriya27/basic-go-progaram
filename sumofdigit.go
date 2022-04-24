package main

import "fmt"

func sumOfdigit(num int) int {
	var remaindor, sum int
	sum = 0

	for num != 0 {
		remaindor = num % 10
		sum = sum + remaindor
		num = num / 10
	}

	return sum
}

func main() {
	var num, sum int
	fmt.Println("enter the number")
	fmt.Scanln(&num)
	sum = sumOfdigit(num)
	fmt.Println(num, "sum Of the number is", sum)

}
