package ffprobe

import (
	"math"
	"strconv"
	"strings"
)

type Fraction string // typically: 25/1, 30000/1001 or 1/90000 etc

// Value returns the float value for the given framerate, or math.NaN() if the value could not be parsed
func (f Fraction) Value() float64 {
	p := strings.IndexByte(string(f), '/')
	if p == -1 {
		// parse as float
		v, err := strconv.ParseFloat(string(f), 64)
		if err != nil {
			return math.NaN()
		}
		return v
	}
	a := string(f[:p])   // 25
	b := string(f[p+1:]) // 1

	ai, err := strconv.ParseUint(a, 10, 64)
	if err != nil {
		return math.NaN()
	}
	bi, err := strconv.ParseUint(b, 10, 64)
	if err != nil {
		return math.NaN()
	}

	return float64(ai) / float64(bi)
}
