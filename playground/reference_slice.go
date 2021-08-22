package main

import "fmt"

func main() {
	// passing a slice to a func and append one element
	s1 := []int{}
	add(s1, 1)
	fmt.Println(s1)

	// passing a slice (created with make) to a func and append one element
	s2 := make([]int, 0)
	add(s2, 1)
	fmt.Println(s2)

	// passing a pointer slice to a func and append one element
	s3 := []int{}
	addR(&s3, 1)
	fmt.Println(s3)

	// passing a pointer slice (created with make) to a func and append one element
	s4 := make([]int, 0)
	addR(&s4, 1)
	fmt.Println(s4)
}

func add(s []int, i int) {
	s = append(s, i)
}

func addR(s *[]int, i int) {
	*s = append(*s, i)
}
