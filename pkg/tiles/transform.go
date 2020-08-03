package tiles

import (
	"kokos_quest/pkg/units"

	"github.com/hajimehoshi/ebiten"
)

type TileTransform int

const (
	TransformEq TileTransform = iota
	Transform90
	Transform180
	Transform270
)

func (k TileTransform) String() string {
	switch k {
	case TransformEq:
		return "TransformEq"
	case Transform90:
		return "Transform90"
	case Transform180:
		return "Transform180"
	case Transform270:
		return "Transform270"
	default:
		return "-- unknown --"
	}
}

func (t *Tile) TransformOp(op *ebiten.DrawImageOptions) {
	w, h := t.tileImage.Size()
	switch t.transform {
	case Transform90:
		op.GeoM.Rotate(units.ToRadians(90))
		op.GeoM.Translate(float64(w), 0)
	case Transform180:
		op.GeoM.Rotate(units.ToRadians(180))
		op.GeoM.Translate(float64(w), float64(h))
	case Transform270:
		op.GeoM.Rotate(units.ToRadians(270))
		op.GeoM.Translate(0, float64(h))
	}
	if t.xMirrored {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(w), 0)
	}
	if t.yMirrored {
		op.GeoM.Scale(1, -1)
		op.GeoM.Translate(0, float64(h))
	}
}
