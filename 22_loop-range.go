package main

import "fmt"

func main() {
	for i, ch := range "Japan 日本" {
		fmt.Printf("%d:%q ", i, ch)
	}
}
