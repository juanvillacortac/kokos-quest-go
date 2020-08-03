package game

import (
	"kokos_quest/pkg/audio"
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/levels"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type TitleScene struct {
	initializedScene bool

	grid map[units.Vec]tiles.Tile
}

func NewTitleScene() *TitleScene {
	return &TitleScene{
		grid: levels.LoadLiteral("level_menu").GetGridTiles(),
	}
}

func (s *TitleScene) InitializedOnce() bool {
	return s.initializedScene
}

func (s *TitleScene) Init() {
	s.initializedScene = true
}

func (s *TitleScene) Update(g *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		audio.PlaySE("Swing")
		g.SetScene(NewGameScene())
	}
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(constants.BackgroundColor)

	units.IterateVecs(constants.TilesPerWidth, constants.TilesPerHeight, func(pos units.Vec) {
		s.grid[pos].Draw(screen, pos)
	})
}
