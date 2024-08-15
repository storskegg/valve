package noise

import (
	"github.com/storskegg/valve/constants"
	"math"
)

// VoltsJohnson calculates the Johnson noise voltage given temperature
// (t in Kelvin), resistance (r in Ohms), and bandwidth (b in Hz). Johnson
// noise is present in any resistor above absolute zero.
func VoltsJohnson(t, r, b float64) float64 {
	return math.Sqrt(4 * constants.Kb * t * r * b)
}

// AmpsJohnson calculates the Johnson noise current given temperature
// (t in Kelvin), resistance (r in Ohms), and bandwidth (b in Hz). Johnson
// noise is present in any resistor above absolute zero.
func AmpsJohnson(t, r, b float64) float64 {
	return math.Sqrt(4 * constants.Kb * t * (1 / r) * b)
}
