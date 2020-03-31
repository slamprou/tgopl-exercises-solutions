package lengthconv

import "fmt"

// Meters type
type Meters float64

// Feet type
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%g m.", m) }
func (f Feet) String() string   { return fmt.Sprintf("%g ft.", f) }

// FToM converts Feet to Meters.
func FToM(f Feet) Meters { return Meters(f * 0.3048) }

// MToF converts Meters to Feet.
func MToF(m Meters) Feet { return Feet(m / 0.3048) }
