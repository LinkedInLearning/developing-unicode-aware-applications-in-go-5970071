package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func main() {
	in, err := os.Open("holmes.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer in.Close()

	io.CopyN(os.Stdout, in, 100)

	fmt.Println("------")
	in.Seek(0, io.SeekStart)
	enc16, _ := charset.Lookup("UTF-16")
	r := enc16.NewDecoder().Reader(in)
	io.CopyN(os.Stdout, r, 100)

	in.Seek(0, io.SeekStart)
	out, err := os.Create("holmes-latin1.text")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer out.Close()

	latin1, _ := charset.Lookup("LATIN1")
	w := latin1.NewEncoder().Writer(out)
	if _, err := io.Copy(w, r); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
