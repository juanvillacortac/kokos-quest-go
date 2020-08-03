package editor

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type wheelDirection int

const (
	wheel_up wheelDirection = iota
	wheel_down
)

func (s *EditorScene) Input() {
	if s.tileInCursor != nil {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			if isAlternativeKey() {
				resetTileTransform(s.tileInCursor)
			} else {
				rotateTile(s.tileInCursor)
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			if isAlternativeKey() {
				mirrorYTile(s.tileInCursor)
			} else {
				mirrorXTile(s.tileInCursor)
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			s.showInfo = !s.showInfo
		}
	}

	switch {
	case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
		if isAlternativeKey() {
			s.chooseTile()
		} else {
			s.placeTile()
		}
	case ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight):
		s.removeTile()
	}

	switch {
	case isWheelMoving(wheel_up):
		if isAlternativeKey() {
			s.incrementKeyIndex()
		} else {
			s.incrementSourcePos()
		}
	case isWheelMoving(wheel_down):
		if isAlternativeKey() {
			s.decrementKeyIndex()
		} else {
			s.decrementSourcePos()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			s.clearGrid()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		s.saveLevel(OutputPath)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		s.setPlayerPos()
	}
}

func isWheelMoving(direction wheelDirection) bool {
	_, y := ebiten.Wheel()
	switch direction {
	case wheel_up:
		return y > 0
	case wheel_down:
		return y < 0
	default:
		return false
	}
}

func isAlternativeKey() bool {
	return ebiten.IsKeyPressed(ebiten.KeyShift)
}
