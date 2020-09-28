package game

import (
	"kokos_quest/pkg/audio"
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/input"
	"kokos_quest/pkg/levels"
	"kokos_quest/pkg/world"

	"github.com/hajimehoshi/ebiten"
)

type GameScene struct {
	initializedScene bool

	w *world.World

	currentLevel uint
}

func NewGameScene() *GameScene {
	return &GameScene{
		w:            levels.Load(1),
		currentLevel: 1,
	}
}

func (s *GameScene) InitializedOnce() bool {
	return s.initializedScene
}

func (s *GameScene) Init() {
	// audio.PlayBGM("Main")
	s.initializedScene = true
}

func (s *GameScene) Update(g *Game, dt float64) {
	if input.IsKeyJustPressed(input.B) {
		audio.PlaySE("Explosion")
	}

	if input.IsKeyJustPressed(input.START) {
		s.w = levels.Load(s.currentLevel)
	}

	s.w.Update(dt)
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(constants.BackgroundColor)
	s.w.Draw(screen)
}
