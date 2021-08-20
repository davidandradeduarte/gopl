package main

import (
	"fmt"
	"os"
)

func main() {
	sep := ""
	for i, a := range os.Args {
		fmt.Printf("%sindex: %d, arg: %s", sep, i, a)
		sep = "\n"
	}
}
