package utils

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func GetRoundedVector(vec3 mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{getRoundedValue(vec3.X()), getRoundedValue(vec3.Y()), getRoundedValue(vec3.Z())}
}

func getRoundedValue(value float64) float64 {
	return math.Trunc(math.Round(value))
}
