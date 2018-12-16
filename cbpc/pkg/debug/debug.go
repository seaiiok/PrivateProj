package main

import (
	"flag"
	"fmt"
)

func main() {
	var a string
	// var b string
	flag.StringVar(&a, "u", "123", "uers")
	// flag.StringVar(&b, "u", "uu", "uers")
	flag.Parse()
	fmt.Println(a)
}
