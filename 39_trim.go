package main

import (
    "fmt"
    "strings"
)

func main() {
	fmt.Println(strings.Trim("foo", "fo"))
	fmt.Println(strings.Trim("bar", "br"))
}
