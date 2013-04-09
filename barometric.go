// Package barometric provides functions for computing the air pressure
// (pascals) or density (kg/m^3) at a given altitude using the barometric
// formula, see: http://en.wikipedia.org/wiki/Barometric_formula
package barometric

import "math"

const (
	WAT_PRESSURE = 0
	WAT_DENSITY  = 1

	G = 9.80665   // Gravitational constant (m/s/s)
	R = 8.31432   // Universal gas constant (N*m/(mol*K))
	M = 0.0289644 // Molar mass of earth's air (kg/mol)

	GMR = G * M / R
)

type empty interface{}

func calc(wat int, height float64) float64 {
	var hb, // Standard height
		lb, // Standard lapse rate
		rb, // Standard density
		pb, // Standard pressure
		tb, // Standard temperature
		c, // interim value (standard density or pressure)
		m, // interim value (mantissa of result)
		e float64 // interim value (exponent of result)

	switch {
	case height < 11000:
		hb, lb, pb, rb, tb = 0, -0.0065, 101325.00, 1.2250, 288.15
	case height < 20000:
		hb, lb, pb, rb, tb = 11000, 0.0, 22632.10, 0.36391, 216.65
	case height < 32000:
		hb, lb, pb, rb, tb = 20000, 0.001, 5474.89, 0.08803, 216.65
	case height < 47000:
		hb, lb, pb, rb, tb = 32000, 0.0028, 868.02, 0.01322, 228.65
	case height < 51000:
		hb, lb, pb, rb, tb = 47000, 0.0, 110.91, 0.00143, 270.65
	case height < 71000:
		hb, lb, pb, rb, tb = 51000, -0.0028, 66.94, 0.00086, 270.65
	default:
		hb, lb, pb, rb, tb = 71000, -0.002, 3.96, 0.000064, 214.65
	}

	if wat == WAT_PRESSURE {
		c = pb
	} else {
		c = rb
	}

	if lb == 0 {
		m = math.E
		e = -GMR * (height - hb) / tb
	} else if wat == WAT_PRESSURE {
		m = tb / (tb + lb*(height-hb))
		e = GMR / lb
	} else if wat == WAT_DENSITY {
		m = (tb + lb*(height-hb)) / tb
		e = -GMR/lb - 1
	}

	return c * math.Pow(m, e)
}

// Get air pressure (pascals) at the given height (h) in meters.
func Pressure(height float64) float64 {
	return calc(WAT_PRESSURE, height)
}

// Get air density (kg/m^3) at the given height (h) in meters.
func Density(height float64) float64 {
	return calc(WAT_DENSITY, height)
}
