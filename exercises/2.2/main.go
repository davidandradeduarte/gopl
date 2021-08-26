package main

import (
	"bufio"
	"fmt"
	"gopl/exercises/2.2/unit-conversion/lengthconv"
	"gopl/exercises/2.2/unit-conversion/tempconv"
	"gopl/exercises/2.2/unit-conversion/weightconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		convert(os.Args[1:])
	} else {
		file := os.Stdin
		fi, err := file.Stat()
		if err != nil {
			panic(err)
		}
		size := fi.Size()
		if size > 0 {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			in := scanner.Text()
			wh := strings.Split(in, " ")
			convert(wh)
		} else {
			fmt.Println("please provide arguments or STDIN to the program")
			os.Exit(0)
		}
	}
}

func convert(numbers []string) {
	for _, ns := range numbers {
		if n, err := strconv.ParseFloat(ns, 64); err == nil {
			f := tempconv.Fahrenheit(n)
			c := tempconv.Celsius(n)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
			m := lengthconv.Meter(n)
			ft := lengthconv.Feet(n)
			fmt.Printf("%s = %s, %s = %s\n",
				m, lengthconv.MToF(m), ft, lengthconv.FToM(ft))
			p := weightconv.Pound(n)
			kg := weightconv.Kilogram(n)
			fmt.Printf("%s = %s, %s = %s\n",
				p, weightconv.PToK(p), kg, weightconv.KToP(kg))
		} else {
			fmt.Printf("%s is not convertible to float64. skipping\n", ns)
			continue
		}
	}
}
