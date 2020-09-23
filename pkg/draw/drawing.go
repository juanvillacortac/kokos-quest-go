package draw

import (
	"image/color"
	"kokos_quest/pkg/global"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var cursorImg *ebiten.Image

func init() {
	cursorImg = getCursorImage()
}

func getCursorImage() *ebiten.Image {
	cursorImg, _ = ebiten.NewImage(3, 3, ebiten.FilterDefault)
	cursorImg.Fill(color.White)
	cursorImg.Set(0, 0, color.Alpha{})
	cursorImg.Set(0, 2, color.Alpha{})
	cursorImg.Set(2, 0, color.Alpha{})
	cursorImg.Set(2, 2, color.Alpha{})
	return cursorImg
}

// Draw an image to screen, specifying the coordinates and an optional `DrawImageOptions`
func DrawWithOp(screen, image *ebiten.Image, px, py int, op *ebiten.DrawImageOptions) {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}
	op.GeoM.Translate(float64(px), float64(py))
	screen.DrawImage(image, op)
}

// Shorthand to `DrawWithOp`, but without `DrawImageOptions` specification
func Draw(screen, image *ebiten.Image, px, py int) {
	DrawWithOp(screen, image, px, py, nil)
}

func DrawWithShadow(screen *ebiten.Image, image *ebiten.Image, px, py int, op *ebiten.DrawImageOptions) {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}
	shadowOp := *op
	shadowOp.ColorM.Scale(0, 0, 0, 0.8)
	DrawWithOp(screen, image, px+1, py+1, &shadowOp)
	DrawWithOp(screen, image, px, py, op)
}

// Draw a cross cursor in the mouse coordinates
func DrawCursor(screen *ebiten.Image) {
	x, y := global.CursorPosition().X, global.CursorPosition().Y

	// Get color in current pixel
	r, g, b, _ := screen.At(x, y).RGBA()
	c := color.RGBA{uint8(r), uint8(g), uint8(b), 255}

	op := &ebiten.DrawImageOptions{}

	// Center to cursor position
	op.GeoM.Translate(-1, -1)

	// Apply color mask in cursor image
	rr := float64(c.R) / 0xff
	gg := float64(c.G) / 0xff
	bb := float64(c.B) / 0xff
	op.ColorM.Scale(0, 0, 0, 1)
	op.ColorM.Translate(rr, gg, bb, 0)

	// Invert color mask
	op.ColorM.Scale(-1, -1, -1, 1)
	op.ColorM.Translate(1, 1, 1, 0)

	DrawWithOp(screen, cursorImg, x, y, op)
}

// Draw a rect given by resolv.
func DrawRect(screen *ebiten.Image, rect *resolv.Rectangle, color color.Color) {
	x, y := float64(rect.X)+1, float64(rect.Y)
	w, h := float64(rect.W)-1, float64(rect.H)-1
	ebitenutil.DrawLine(screen, x, y, x, y+h+1, color)
	ebitenutil.DrawLine(screen, x+w, y, x+w, y+h, color)

	ebitenutil.DrawLine(screen, x, y, x+w, y, color)
	ebitenutil.DrawLine(screen, x, y+h, x+w, y+h, color)
}
