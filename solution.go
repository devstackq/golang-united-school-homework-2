package square

import "math"

type Number float64

const (
	SidesTriangle Number = 3
	SidesSquare   Number = 4
	SidesCircle   Number = 0
)

func CalcSquare(sideLen float64, sidesNum Number) float64 {
	if sidesNum == SidesTriangle {
		return math.Sqrt(float64(sidesNum)) * (float64(sideLen) * float64(sideLen)) / 4
	}
	if sidesNum == SidesSquare {
		return math.Pi * sideLen * sideLen
	}
	if sidesNum == SidesCircle {
		return sideLen * float64(SidesSquare)
	}
	return 0
}
