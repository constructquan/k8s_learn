package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	name := flag.String("name", "gogo", "something you want to say!")
	fmt.Println(os.Args)
	
	flag.Parse() //这一步挺容易忽视的。
	if len(os.Args) > 1 {
		fmt.Printf("hello world ... %s\n", *name)
	}else {
		os.Exit(-1)
	}
}
