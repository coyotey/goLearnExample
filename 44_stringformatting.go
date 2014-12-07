package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", p)
	fmt.Printf("%d\n", p)

	fmt.Printf("%b\n", p)
	fmt.Printf("%c\n", p)

	fmt.Printf("%x\n", p)
	fmt.Printf("%f\n", p)

	fmt.Printf("%e\n", 12340000000.0)
	fmt.Printf("%E\n", 12340000000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
}
