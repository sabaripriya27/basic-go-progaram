package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sample := "college friends"

	fmt.Printf("Length of given string is %d\n", utf8.RuneCountInString(sample))
}
