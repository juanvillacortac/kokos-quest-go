package world

import (
	"image/color"

	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/global"
	"kokos_quest/pkg/player"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type World struct {
	mapDims units.Dims
	grid    map[units.Vec]tiles.Tile
	player  *player.Player
	space   *resolv.Space
}

func initSpace(playerPos units.Vec, grid map[units.Vec]tiles.Tile) *resolv.Space {
	space := resolv.NewSpace()
	space.Clear()

	playerPos = units.TileToVec(playerPos)

	for pos, tile := range grid {
		pos = pos.ToAbsolute()
		rect := resolv.NewRectangle(
			int32(pos.X), int32(pos.Y),
			units.TileSize, units.TileSize,
		)
		rect.SetData(tile)
		rect.AddTags("tile")
		space.Add(rect)
	}

	return space
}

var isColliding bool

func (w *World) Update(dt float64) {
	w.player.Update(w.space, dt)

	isColliding = w.space.WouldBeColliding(w.player.Rect(), 0, 0)
}

func (w *World) Draw(screen *ebiten.Image) {
	w.player.Draw(screen)
	w.drawTiles(screen)

	if global.Config.Debug {
		w.drawDebug(screen)
	}
}

func (w *World) drawDebug(screen *ebiten.Image) {
	collide := "No colliding"
	if isColliding {
		collide = "Colliding!"
	}

	for _, shape := range *w.space.FilterByTags("tile") {
		rect, ok := shape.(*resolv.Rectangle)
		if !ok {
			break
		}
		c := color.RGBA{60, 255, 60, 255}
		if w.space.IsColliding(shape) {
			draw.DrawRect(screen, rect, c)
		}
	}

	draw.DrawRect(screen, w.player.Rect(), color.White)

	ebitenutil.DebugPrintAt(screen, collide, 0, 16)
	ebitenutil.DebugPrintAt(screen, w.player.Pos().String(), 0, 32)
}

func (w *World) drawTiles(screen *ebiten.Image) {
	units.IterateVecs(w.mapDims.W, w.mapDims.H, func(pos units.Vec) {
		w.grid[pos].Draw(screen, pos)
	})
}

func (w *World) GetGridTiles() map[units.Vec]tiles.Tile {
	return w.grid
}
