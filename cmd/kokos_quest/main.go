package main

import (
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	w, h, s := constants.ScreenWidth, constants.ScreenHeight, 2
	instance, err := game.NewGame("Koko's Quest", w, h, s, true)
	if err != nil {
		panic(err)
	}
	instance.SetInitialScene(func() game.Scene { return game.NewGameScene() })
	if err := ebiten.RunGame(instance); err != nil {
		panic(err)
	}
}
