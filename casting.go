package sprint

import "math"

func Casting(n float32) int {
	return int(math.Round(float64(n)))
}
