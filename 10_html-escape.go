package main

import (
	"fmt"
	"html"
)

func main() {
	fmt.Println(html.EscapeString("<>"))
}
