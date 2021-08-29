package main

import "fmt"

func main() {
	x := 1e100 // integer literal
	y := 0xf   // integer literal
	fmt.Println(x, int(x))
	fmt.Println(y, fmt.Sprintf("%016b", y)) // print binary value with 16 bits

	z := 0x000f
	fmt.Printf("%d %#[1]x %[1]X %#08[1]b", z)
}
