package main

import "fmt"

func main() {
	sample := "abcd"
	r := []rune(sample)

	fmt.Printf("Before %s\n", string(r))
	r[0], r[3] = r[2], r[3]

	fmt.Printf("After %s\n", string(r))

	a := "123"
	b := "xyz"
	fmt.Printf("Before a:%s b:%s\n", a, b)
	a, b = b, a
	fmt.Printf("After a:%s b:%s\n", a, b)
}
