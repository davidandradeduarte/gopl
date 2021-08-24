package main

import "fmt"

func main() {
	var a *int    // pointer to nothing, zero value: nil
	b := new(int) // pointer (*int) to unamed variable with int zero value: 0
	fmt.Println(a)
	// fmt.Println(*a) // panic: invalid memory address or nil pointer dereference
	fmt.Println(b)
	fmt.Println(*b)

	fmt.Println("-")

	x := 1
	z := &x         // z pointer to x
	*z++            // increment the variable that points to
	fmt.Println(x)  // x value
	fmt.Println(*z) // z, value of the variable that points to (x)
	// different memory addresses tho
	fmt.Println(&x)
	fmt.Println(&z)

	fmt.Println("-")

	var t *struct{}
	o := new(struct{})
	p := new(struct{})
	fmt.Println(t)
	fmt.Println(o) // memory address is the same as p (&{}) because struct{} is size zero and carries no information
	fmt.Println(p) // memory address is the same as o (&{}) because struct{} is size zero and carries no information
	// fmt.Println(*t) // panic: invalid memory address or nil pointer dereference
	fmt.Println(*o)
	fmt.Println(*p)
}
