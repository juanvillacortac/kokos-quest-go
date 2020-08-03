// +build !js

package util

import "github.com/hajimehoshi/ebiten"

func IsFullscreen() bool {
	return ebiten.IsFullscreen()
}

func SetFullscreen(fullscreen bool) {
	ebiten.SetFullscreen(fullscreen)
}
