package units

import (
	"fmt"
	"math"
)

const TileSize = 16

type Dims struct {
	W, H int
}

func (d Dims) GetWH() (int, int) {
	return d.W, d.H
}

func (d Dims) String() string {
	return fmt.Sprintf("{W: %3d, H: %3d}", d.W, d.H)
}

type Vec struct {
	X, Y int
}

type VecFloat struct {
	X, Y float64
}

func IterateVecs(width, height int, callback func(Vec)) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := Vec{X: x, Y: y}
			callback(v)
		}
	}
}

func (v Vec) GetXY() (int, int) {
	return v.X, v.Y
}

func (v Vec) ToAbsolute() Vec {
	return TileToVec(v)
}

func TileToVec(pos Vec) Vec {
	convert := func(t int) int {
		return t * TileSize
	}
	return Vec{X: convert(pos.X), Y: convert(pos.Y)}
}

func ToRadians(angle float64) float64 {
	return angle * math.Pi / 180
}

func (v Vec) String() string {
	return fmt.Sprintf("{X: %3d, Y: %3d}", v.X, v.Y)
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
