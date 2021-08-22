package main

import (
	"encoding/json"
	"fmt"
)

// apparently 
func main() {
	// passing a map to a func and append one element
	m1 := map[string]int{}
	add(m1, "1", 1)
	j1, _ := json.Marshal(m1)
	fmt.Println(string(j1))

	// passing a map (created with make) to a func and append one element
	m2 := make(map[string]int)
	add(m2, "2", 2)
	j2, _ := json.Marshal(m2)
	fmt.Println(string(j2))
}

func add(m map[string]int, s string, i int) {
	m[s] = i
}
