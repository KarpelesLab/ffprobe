package ffprobe

import (
	"math"
	"strconv"
	"strings"
)

type Fraction string // typically: 25/1, 30000/1001 or 1/90000 etc

// Value returns the float value for the given framerate, or math.NaN() if the value could not be parsed
func (f Fraction) Value() float64 {
	a, b := f.Frac()
	if b == 0 {
		return math.NaN()
	}
	return float64(a) / float64(b)
}

func (f Fraction) Frac() (int, int) {
	p := strings.IndexByte(string(f), '/')
	if p == -1 {
		// parse as int
		v, err := strconv.ParseInt(string(f), 10, 64)
		if err != nil {
			return 0, 0
		}
		return int(v), 1
	}
	a := string(f[:p])   // 25
	b := string(f[p+1:]) // 1

	ai, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return 0, 0
	}
	bi, err := strconv.ParseUint(b, 10, 64)
	if err != nil {
		return 0, 0
	}

	return int(ai), int(bi)
}
