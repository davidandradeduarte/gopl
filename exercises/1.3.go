package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// add 100 elements to os.Args
	for i := 0; i < 100; i++ {
		os.Args = append(os.Args, fmt.Sprint(i))
	}
	now1 := time.Now()
	echo1()
	fmt.Println(time.Since(now1)) // ~ 53.485µs

	now2 := time.Now()
	echo2()
	fmt.Println(time.Since(now2)) // ~ 10.401µs

	now3 := time.Now()
	echo3()
	fmt.Println(time.Since(now3)) // ~ 3.318µs
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
