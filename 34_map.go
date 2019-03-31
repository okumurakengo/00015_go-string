package main

import (
    "fmt"
    "strings"
)

func main() {
	f := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f, "あい"))
}
