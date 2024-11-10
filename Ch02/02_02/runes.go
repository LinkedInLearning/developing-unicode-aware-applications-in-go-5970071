package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "I ♡ Go"

	fmt.Println("len:", len(text)) // 8 (not 6)
	fmt.Printf("text[0]: %c (%T)\n", text[0], text[0])
	fmt.Printf("text[2]: %c\n", text[2])

	for i, c := range text {
		fmt.Printf("%d: %c (%T)\n", i, c, c)
	}

	fmt.Println("len (runes):", utf8.RuneCountInString(text))

	rs := []rune(text)
	fmt.Println("len ([]rune):", len(rs))
	fmt.Printf("rs[2]: %c\n", rs[2])
}
