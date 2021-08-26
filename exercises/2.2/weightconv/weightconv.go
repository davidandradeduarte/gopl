package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (c Pound) String() string    { return fmt.Sprintf("%glb", c) }
func (f Kilogram) String() string { return fmt.Sprintf("%gkg", f) }

// PToK converts a Pound weight to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KToP converts a Kilogram weight to Pound.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }
