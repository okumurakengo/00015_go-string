package main

import "fmt"

func main() {
	s := "Japan 日本"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%q ", s[i])
	}
}
