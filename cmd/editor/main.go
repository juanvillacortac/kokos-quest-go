package main

import (
	"fmt"
	"os"

	"kokos_quest/cmd/editor/internal"
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	if len(os.Args) > 1 {
		editor.OutputPath = os.Args[1]
	} else {
		fmt.Printf("USAGE: \n\t%s [FILE]\n", os.Args[0])
		return
	}

	w, h, s := constants.ScreenWidth, constants.ScreenHeight, 2
	instance, err := game.NewGame("Koko's Quest Map Editor", w, h, s, false)
	if err != nil {
		panic(err)
	}
	instance.SetInitialScene(func() game.Scene { return editor.NewEditorScene() })
	if err := ebiten.RunGame(instance); err != nil {
		panic(err)
	}
}
