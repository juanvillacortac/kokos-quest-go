package constants

import "image/color"

const (
	TileSize = 16

	TilesPerWidth  = 20
	TilesPerHeight = 20

	ScreenWidth  = TileSize * TilesPerWidth
	ScreenHeight = TileSize * TilesPerHeight
)

var (
	BackgroundColor color.RGBA = color.RGBA{24, 20, 37, 255}
)
