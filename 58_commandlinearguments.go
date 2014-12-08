package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithPro := os.Args
	argsWithoutPro := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithPro)
	fmt.Println(argsWithoutPro)
	fmt.Println(arg)
}
