package units

import (
	"fmt"
	"kokos_quest/pkg/constants"
	"math"
)

type Dims struct {
	W, H int
}

type Vec struct {
	X, Y int
}

func IterateVecs(width, height int, callback func(Vec)) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := Vec{X: x, Y: y}
			callback(v)
		}
	}
}

func (v Vec) ToAbsolute() Vec {
	return TileToVec(v)
}

func (v Vec) String() string {
	return fmt.Sprintf("{X: %3d, Y: %3d}", v.X, v.Y)
}

func TileToVec(pos Vec) Vec {
	convert := func(t int) int {
		return t * constants.TileSize
	}
	return Vec{X: convert(pos.X), Y: convert(pos.Y)}
}

func ToRadians(angle float64) float64 {
	return angle * math.Pi / 180
}

type HorizDir int

const (
	Left HorizDir = iota
	Right
)

type VerDir int

const (
	Up VerDir = iota
	Down
)
