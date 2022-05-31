package helpers

import "math"

func DistanceBetween(x1, y1, x2, y2 float64) float64 {
	y := math.Pow(y2-y1, 2.0)
	x := math.Pow(x2-x1, 2)
	results := math.Sqrt(y + x)
	return results
}

func Colision(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		h1+y1 > y2
}
