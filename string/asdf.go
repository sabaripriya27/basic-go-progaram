package main

import (
	"fmt"
	"strings"
)

func main() {

	sample := " This is a sample string"
	noSpaceString := strings.ReplaceAll(sample, " is", "")
	fmt.Println(noSpaceString)
}
