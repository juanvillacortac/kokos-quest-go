package editor

import (
	"image/color"
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/units"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	grayTile        *ebiten.Image
	whiteTile       *ebiten.Image
	placeholderTile *ebiten.Image
	cursorTile      *ebiten.Image
	playerTile      *ebiten.Image
)

func init() {
	whiteTile, _ = ebiten.NewImage(constants.TileSize, constants.TileSize, ebiten.FilterDefault)
	whiteTile.Fill(constants.BackgroundColor)

	grayTile, _ = ebiten.NewImage(constants.TileSize, constants.TileSize, ebiten.FilterDefault)
	r, g, b, _ := constants.BackgroundColor.RGBA()
	grayC := color.RGBA{uint8(r + 20), uint8(g + 20), uint8(b + 20), 255}
	grayTile.Fill(grayC)

	placeholderTile, _ = ebiten.NewImage(constants.TileSize, constants.TileSize, ebiten.FilterDefault)
	placeholderC := color.RGBA{20, 20, 255, 60}
	placeholderTile.Fill(placeholderC)

	cursorTile, _ = ebiten.NewImage(constants.TileSize, constants.TileSize, ebiten.FilterDefault)
	cursorC := color.RGBA{20, 255, 20, 60}
	cursorTile.Fill(cursorC)

	playerTile, _ = ebiten.NewImage(constants.TileSize, constants.TileSize, ebiten.FilterDefault)
	playerC := color.RGBA{255, 20, 20, 60}
	playerTile.Fill(playerC)
}

func (s *EditorScene) DrawTileCursor(screen *ebiten.Image) {
	if s.tileInCursor != nil {
		op := &ebiten.DrawImageOptions{}
		op.ColorM.Scale(1, 1, 1, 0.5)
		draw.Draw(screen, cursorTile, s.cursorPos.X*constants.TileSize, s.cursorPos.Y*constants.TileSize)
		s.tileInCursor.Build().DrawWithOp(screen, s.cursorPos, op)
	}
}

func (s *EditorScene) DrawTiles(screen *ebiten.Image) {
	units.IterateVecs(constants.TilesPerWidth, constants.TilesPerHeight, func(pos units.Vec) {
		tb, ok := s.grid[pos]
		if ok {
			posAbs := pos.ToAbsolute()
			draw.Draw(screen, placeholderTile, posAbs.X, posAbs.Y)
			tb.Build().Draw(screen, pos)
		}
	})
}

func (s *EditorScene) DrawPlayerPos(screen *ebiten.Image) {
	if s.playerPos == nil {
		return
	}
	x, y := s.playerPos.X*constants.TileSize, s.playerPos.Y*constants.TileSize
	draw.Draw(screen, playerTile, x, y)
	ebitenutil.DebugPrintAt(screen, "P", x+5, y)
}

func DrawBackground(screen *ebiten.Image) {
	sw, sh := draw.CountTiledImage(screen, constants.TileSize)
	for y := 0; y < sh; y++ {
		for x := 0; x < sw; x++ {
			ts := constants.TileSize
			dx, dy := float64(x*ts), float64(y*ts)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(dx, dy)
			if y%2 == 0 {
				if x%2 != 0 {
					screen.DrawImage(whiteTile, op)
				} else {
					screen.DrawImage(grayTile, op)
				}
			} else {
				if x%2 == 0 {
					screen.DrawImage(whiteTile, op)
				} else {
					screen.DrawImage(grayTile, op)
				}
			}
		}
	}
}
