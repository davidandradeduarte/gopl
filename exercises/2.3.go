package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	t := time.Now()
	PopCount(10)
	fmt.Println(time.Since(t))
	t = time.Now()
	PopCountL(10)
	fmt.Println(time.Since(t))
}

// PopCountL returns the population count (number of set bits) of x.
func PopCountL(x uint64) int { // 173ns
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int { // 89ns
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
