package main

import (
	"fmt"
	"strings"
)

func main() {
	//Input string contains the prefix
	res := strings.HasPrefix("abcdef", "ab")
	fmt.Println(res)

	//Input string doesn't contain the prefix
	res = strings.HasPrefix("abcdef", "ac")
	fmt.Println(res)
	//Input string contains the suffix
	res = strings.HasSuffix("abcdef", "ef")
	fmt.Println(res)

	//Input string doesn't contain the suffix
	res = strings.HasPrefix("abcdef", "fg")
	fmt.Println(res)
}
