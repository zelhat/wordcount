package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var size int

	if len(os.Args) > 1 {

		str := os.Args[1]

		if len(str) != 0 {
			size = len(strings.Split(str, " "))
		}

	}

	fmt.Println(size)
}
