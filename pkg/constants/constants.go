package constants

import (
	"image/color"
	"kokos_quest/pkg/units"
)

const (
	TilesPerWidth  = 20
	TilesPerHeight = 20

	ScreenWidth  = units.TileSize * TilesPerWidth
	ScreenHeight = units.TileSize * TilesPerHeight
)

var (
	BackgroundColor color.RGBA = color.RGBA{24, 20, 37, 255}
)
