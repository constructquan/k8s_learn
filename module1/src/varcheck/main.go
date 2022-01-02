package main

import (
	"fmt"
)

func main() {
	var (
		a  int     = 10
		b  int8    = 54
		s  string  = "a"
		c  rune    = 'A'
		t  bool    = true
		f1 float32 = 16777216.0
		f2 float32 = 16777217.0
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(c)
	fmt.Println(t)
	if f1 == f2 {
		fmt.Print("true")
	}
}
