package main

import (
	"fmt"
	"strings"
)

func main() {
	//It will replaces 1 instance of a with 1
	res := strings.Replace("abcdabxyabr", "a", "1", 1)
	fmt.Println(res)

	//It will replace all instances of a with 1
	res = strings.Replace("abcdabxyabr", "a", "1", -1)
	fmt.Println(res)

	res = strings.ReplaceAll("abcdabxyabr", "ab", "12")
	fmt.Println(res)

	res = strings.ReplaceAll("abcdabxyabr", "", "12")
	fmt.Println(res)

	//Output will be 1 as "bc" is present in "abcdef" at index 1
	r := strings.LastIndex("abcdef", "bc")
	fmt.Println(r)

	//Output will be 8 as "cd" is present in "abcdefabcdef" at index 8
	r = strings.LastIndex("abcdefabcdef", "cd")
	fmt.Println(r)

	//Output will be -1 as "ba" is not present in "abcdef"
	r = strings.LastIndex("abcdef", "ba")
	fmt.Println(r)
	//delete
	res = strings.ReplaceAll("abcdabxyabr", "a", "")
	fmt.Println(res)

	res = strings.Repeat("abc ", 4)
	fmt.Println(res)
}
