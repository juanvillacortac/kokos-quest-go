package player

import (
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/physics"
	"kokos_quest/pkg/units"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
)

const (
	PlayerWidth  = units.TileSize
	PlayerHeight = units.TileSize

	RectWidth  = int32(PlayerWidth - 2)
	RectHeight = int32(PlayerHeight)

	RectOffsetX = 1
)

type Player struct {
	dir        units.HorizDir
	pos        units.Vec
	kinematics physics.Kinematics2D

	rect *resolv.Rectangle
}

func NewPlayer(pos units.Vec, dir units.HorizDir) *Player {
	pos = pos.ToAbsolute()
	rect := resolv.NewRectangle(int32(pos.X+1), int32(pos.Y), RectWidth, RectHeight)
	rect.AddTags("player")

	return &Player{
		dir: dir,
		kinematics: physics.MakeKinematics2D(pos, units.VecFloat{}),

		rect: rect,
	}
}

func (p *Player) Pos() units.Vec {
	return p.kinematics.Position()
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

func (p *Player) Update(space *resolv.Space, dt float64) {
	p.updateY()

	p.updateRect()
}

func (p *Player) updateRect() {
	pos := p.Pos()
	x, y := int32(pos.X+RectOffsetX), int32(pos.Y)
	p.rect.SetXY(x, y)
}

func (p *Player) Draw(screen *ebiten.Image) {
	img := draw.GetTiledImage("PlayerStanding", 0, 0)
	op := &ebiten.DrawImageOptions{}
	p.transformOp(op)
	pos := p.Pos()
	draw.DrawWithShadow(screen, img, pos.X, pos.Y, op)
}
