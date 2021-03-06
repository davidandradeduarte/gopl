package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)/1.8 + 273.15) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*1.8 - 459.67) }

func main() {
	fmt.Printf("absolute zero:\n%v\n%v\n%v\n\n", AbsoluteZeroC, CToF(AbsoluteZeroC), CToK(AbsoluteZeroC))
	fmt.Printf("freezing:\n%v\n%v\n%v\n\n", FreezingC, CToF(FreezingC), CToK(FreezingC))
	fmt.Printf("boiling:\n%v\n%v\n%v\n\n", BoilingC, CToF(BoilingC), CToK(BoilingC))
}
