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
	PopCountLoop(10)
	fmt.Println(time.Since(t))

	t = time.Now()
	PopCountShift(10)
	fmt.Println(time.Since(t))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int { // 174ns
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountL returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int { // 136ns
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

func PopCountShift(x uint64) int { // 94ns
	count := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return count
}
