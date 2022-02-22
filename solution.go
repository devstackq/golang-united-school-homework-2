package square

import "math"

type Number int

const (
	SidesTriangle Number = 3
	SidesSquare   Number = 4
	SidesCircle   Number = 0
)

func CalcSquare(sideLen float64, sidesNum Number) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
