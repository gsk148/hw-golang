package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reverseString := stringutil.Reverse("Hello, OTUS!")
	fmt.Print(reverseString)
}
