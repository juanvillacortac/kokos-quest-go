package global

import (
	"kokos_quest/pkg/units"

	"github.com/hajimehoshi/ebiten"
)

var (
	ScreenSize        units.Dims
	VirtualScreenSize units.Dims
	ScaleFactor       float64 = 1
)

func CursorPosition() units.Vec {
	s := int(ScaleFactor)
	x, y := ebiten.CursorPosition()
	return units.Vec{
		X: x / s,
		Y: y / s,
	}
}

type ConfigStruct struct {
	Debug bool

	Shader  bool
	HighDpi bool
}

var Config = &ConfigStruct{}
