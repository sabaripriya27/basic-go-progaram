package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter any string to find ASCII values=")
	strData, _ := reader.ReadString('\n')
	for i := 0; i < len(strData); i++ {
		fmt.Printf("the ASCII values of%c=%d", strData[i], strData[i])
	}

}
