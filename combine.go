package main

import (
	"fmt"
	"math"
)

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
func amstrong(num int) int {

	var temp, sum, start, end, remainder int

	fmt.Print("Enter the Minimum and Maximum limit of Amstrong = ")
	fmt.Scanln(&start, &end)

	fmt.Print("Palindrome Numbers between ", start, " and ", end, " are : ")
	for num = start; num <= end; num++ {
		sum = 0
		for temp = num; temp > 0; temp = temp / 10 {
			remainder = temp % 10
			sum = sum + remainder*remainder*remainder
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
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
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
func calcFactorial(factorialnum int) int {
	factorial := 1
	for i := 1; i <= factorialnum; i++ {
		factorial = factorial * i
	}
	return factorial
}
func perfect(n int) int {

	fmt.Print("Enter any number: ")
	fmt.Scanf("%d", &n)
	sum1 := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum1 += i
		}
	}
	if sum1 == n {
		print("The number is a Perfect number!")
	} else {
		print("The number is not a Perfect number!")
	}
	return sum1
}
func main() {
	var num int
	palindrom(num)
	fmt.Println("")
	fmt.Println("...................Amstrang Number.................")
	var ams int
	amstrong(ams)
	fmt.Println("")
	fmt.Println("....................Prime Number.....................")
	for i := 1; i <= 100; i++ {
		if IsPrime(i) {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("")
	fmt.Println(".....................Sum Of Digit...................")
	var digit, sum int
	fmt.Println("enter the number")
	fmt.Scanln(&digit)
	sum = sumOfdigit(digit)
	fmt.Println(digit, "sum Of the number is", sum)

	fmt.Println("")
	fmt.Println(".....................Reverse Number...................")
	var number, Reverse int
	Reverse = reverse(number)
	fmt.Println("The Rverses Number is", Reverse)
	fmt.Println("")
	fmt.Println(".....................Fibonacci Series...................")
	var series int

	printFibonacciSeries(series)
	fmt.Println("")
	fmt.Println(".....................Factorial Number...................")
	var factorialnum, factorial int

	fmt.Print("Enter any Number to find the Factorial = ")
	fmt.Scanln(&factorialnum)

	factorial = calcFactorial(factorialnum)
	fmt.Println("The Factorial of ", factorialnum, " = ", factorial)
	fmt.Print("")
	fmt.Println("..................Perfect Number.............")
	var pernum int
	perfect(pernum)
}
