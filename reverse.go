package main

import "fmt"

func reverse(num int) int {
	var remaindor, rev int

	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	rev = 0

	for num != 0 {
		remaindor = num % 10
		rev = rev*10 + remaindor
		num = num / 10
		if num == 0 {
			break
		}
	}

	return rev
}

func main() {
	var num, Reverse int
	Reverse = reverse(num)
	fmt.Println("The Rverses Number is", Reverse)

}
