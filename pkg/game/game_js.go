// +build js

package game

import (
	"kokos_quest/pkg/global"
	"syscall/js"

	"github.com/hajimehoshi/ebiten"
)

func (g *Game) InitStuff() {
	doc := js.Global().Get("document")
	screenSizeObj := js.ValueOf(map[string]interface{}{
		"width":  global.ScreenWidth,
		"height": global.ScreenHeight,
		"scale":  ebiten.DeviceScaleFactor(),
	})
	doc.Set("__Game__ScreenSize", screenSizeObj)
}
