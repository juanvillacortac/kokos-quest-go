package tiles

import (
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/units"

	"github.com/hajimehoshi/ebiten"
)

type TileBuilder struct {
	Key     string
	SourceX int
	SourceY int

	Kind      TileKind
	Transform TileTransform
	XMirrored bool
	YMirrored bool
}

func (t TileBuilder) Build() Tile {
	return Tile{
		kind:      t.Kind,
		tileImage: draw.GetTiledImage(t.Key, t.SourceX, t.SourceY),
		transform: t.Transform,
		xMirrored: t.XMirrored,
		yMirrored: t.YMirrored,
	}
}

type Tile struct {
	kind      TileKind
	tileImage *ebiten.Image
	transform TileTransform
	xMirrored bool
	yMirrored bool
}

func (t Tile) Draw(screen *ebiten.Image, pos units.Vec) {
	t.DrawWithOp(screen, pos, nil)
}

func (t Tile) DrawWithOp(screen *ebiten.Image, pos units.Vec, op *ebiten.DrawImageOptions) {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}
	if t.tileImage != nil {
		pos = units.TileToVec(pos)
		t.TransformOp(op)
		draw.DrawWithShadow(screen, t.tileImage, pos.X, pos.Y, op)
	}
}
