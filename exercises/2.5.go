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

	t = time.Now()
	PopCountRightMost(10)
	fmt.Println(time.Since(t))
}

// PopCountL returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int { // 109ns
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int { // 93ns
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShift(x uint64) int { // 91ns
	count := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return count
}

func PopCountRightMost(x uint64) int { // 50ns
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
