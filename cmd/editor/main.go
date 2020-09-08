package main

import (
	"kokos_quest/cmd/editor/internal"
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	w, h, s := constants.ScreenWidth, constants.ScreenHeight, 2
	instance, err := game.NewGame("Koko's Quest Map Editor", w, h, s, false)
	if err != nil {
		panic(err)
	}
	instance.SetInitialScene(func() game.Scene {
		return editor.NewEditorScene(LoadLevelHandler, SaveLevelHandler)
	})
	if err := ebiten.RunGame(instance); err != nil {
		panic(err)
	}
}
