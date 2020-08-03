package player

import (
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/input"
	"kokos_quest/pkg/units"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
)

const (
	PlayerWidth  = constants.TileSize
	PlayerHeight = constants.TileSize

	RectWidth  = int32(PlayerWidth - 2)
	RectHeight = int32(PlayerHeight)

	RectOffsetX = 1
)

type Player struct {
	dir units.HorizDir
	pos units.Vec

	rect *resolv.Rectangle
}

func NewPlayer(pos units.Vec, dir units.HorizDir) *Player {
	pos = pos.ToAbsolute()
	rect := resolv.NewRectangle(int32(pos.X+1), int32(pos.Y), RectWidth, RectHeight)
	rect.AddTags("player")

	return &Player{
		dir: dir,
		pos: pos,

		rect: rect,
	}
}

func (p *Player) Pos() units.Vec {
	return p.pos
}

func (p *Player) Rect() *resolv.Rectangle {
	return p.rect
}

func (p *Player) texAndMirror() (*ebiten.Image, bool) {
	img := draw.GetTiledImage("PlayerStanding", 0, 0)
	return img, p.dir == units.Left
}

func (p *Player) transformOp(op *ebiten.DrawImageOptions) {
	if p.dir == units.Left {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(PlayerWidth), 0)
	}
}

func (p *Player) Update(space *resolv.Space) {
	if input.IsKeyPressed(input.LEFT) {
		p.pos.X -= 1
		if p.dir == units.Right {
			p.dir = units.Left
		}
	} else if input.IsKeyPressed(input.RIGHT) {
		p.pos.X += 1
		if p.dir == units.Left {
			p.dir = units.Right
		}
	}

	if input.IsKeyPressed(input.UP) {
		p.pos.Y -= 1
	} else if input.IsKeyPressed(input.DOWN) {
		p.pos.Y += 1
	}

	if input.IsKeyPressed(input.A) {
		if p.dir == units.Left {
			p.pos.X--
		} else {
			p.pos.X++
		}
	}

	p.updateRect()
}

func (p *Player) updateRect() {
	x, y := int32(p.pos.X+RectOffsetX), int32(p.pos.Y)
	p.rect.SetXY(x, y)
}

func (p *Player) Draw(screen *ebiten.Image) {
	img := draw.GetTiledImage("PlayerStanding", 0, 0)
	op := &ebiten.DrawImageOptions{}
	p.transformOp(op)
	draw.DrawWithShadow(screen, img, p.pos.X, p.pos.Y, op)
}
