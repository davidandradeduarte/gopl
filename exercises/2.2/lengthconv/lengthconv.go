package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (c Meter) String() string { return fmt.Sprintf("%gm", c) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

// MToF converts a Meter length to Feet.
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

// FToM converts a Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f / 3.2808) }
