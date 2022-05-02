package main

import (
	"fmt"
	"strings"
)

func main() {
	//chance uper charcter in lower character
	res := strings.ToLower("ABC")
	fmt.Println(res)

	res = strings.ToLower("ABC12$a")
	fmt.Println(res)
	//chance lower charcter in uper character
	res = strings.ToUpper("abc")
	fmt.Println(res)

	res = strings.ToUpper("abc12$")
	fmt.Println(res)
	//chnce eveery first character in captal letter
	res = strings.Title("this is a test sentence")
	fmt.Println(res)
}
