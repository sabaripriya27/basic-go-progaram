package main

import "fmt"

func swap1(a, b int) {
	var c int
	fmt.Println("enter the two number")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println("Before swapping")
	fmt.Println("a :", a)
	fmt.Println("b :", b)
	c = a
	a = b
	b = c
	fmt.Println("After swapping")
	fmt.Println("a :", a)
	fmt.Println("b :", b)

}
func swap2(a1, b1 int) {

	fmt.Println("enter the two number")

	a1 = 10
	b1 = 20
	fmt.Println("Before swapping")
	fmt.Println("a1 :", a1)
	fmt.Println("b1 :", b1)
	a1 = a1 + b1
	b1 = a1 - b1
	a1 = a1 - b1

	fmt.Println("After swapping")
	fmt.Println("a1 :", a1)
	fmt.Println("b1 :", b1)

}

func swap3(a, b int) {

	fmt.Println("enter the two number")
	a = 100
	b = 200
	fmt.Println("Before swapping")
	fmt.Println("a :", a)
	fmt.Println("b :", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("After swapping")
	fmt.Println("a :", a)
	fmt.Println("b :", b)

}
func main() {
	var a, b int
	fmt.Println("..............Normal Method...........")
	//var a, b int
	swap1(a, b)
	fmt.Println("")
	fmt.Println("...........without using third veriable...........")
	var a1, b1 int
	swap2(a1, b1)
	fmt.Println("")
	fmt.Println("...........By using Xor method...........")
	swap3(a, b)
}
