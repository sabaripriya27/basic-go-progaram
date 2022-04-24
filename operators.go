package main

import "fmt"

/*
func arthamatic(x, y int) int {
	var z int
	fmt.Scanf("%d %d", &x, &y)
	z = x + y

	return z
}*/

func main() {
	var a, b int
	fmt.Println("enter the a and b vlues")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(".............arthamatic Operator.............")
	fmt.Println("adddtion:", a+b)
	fmt.Println("subraction:", a-b)
	fmt.Println("multiplcation:", a*b)
	fmt.Println("divition:", a/b)
	fmt.Println("modular:", a%b)
	a++
	fmt.Println("increment a:", a)
	b--
	fmt.Println("decrement b:", b)
	fmt.Println("..........Bitwise Operator.............")
	fmt.Println("AND :", a&b)
	fmt.Println("OR :", a|b)
	fmt.Println("XOR :", a^b)
	/*fmt.Println("Complement :",~a)
	fmt.Println("complement :",~-a)*/
	fmt.Println("right shift :", (a << 2))
	fmt.Println("left shift :", (b >> 2))
	fmt.Println(".........comparision oprator............")
	fmt.Println("a==b ;", a == b)
	fmt.Println("a!=b:", a != b)
	fmt.Println("a>b :", a > b)
	fmt.Println("a<=b:", a <= b)
	fmt.Println("...........Logical Operator..........")
	fmt.Println("logical AND :", (a > 3 && b < 3))
	fmt.Println("logical OR:", (a < 11 || b > 4))

}
