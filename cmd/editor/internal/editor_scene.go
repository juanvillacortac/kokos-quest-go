package editor

import (
	"fmt"
	"math"

	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/game"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	keys       = draw.Keys()
	OutputPath string
)

type EditorScene struct {
	initializedScene bool

	keyIndex  int
	sourcePos units.Vec

	cursorPos    units.Vec
	tileInCursor *tiles.TileBuilder
	grid         map[units.Vec]tiles.TileBuilder
	playerPos    *units.Vec

	showInfo bool
	saved    bool
}

func NewEditorScene() *EditorScene {
	level, err := loadLevel(OutputPath)
	if err != nil {
		panic(err)
	}
	var playerPos *units.Vec
	if level.PlayerPos.X != -1 {
		playerPos = &level.PlayerPos
	}
	return &EditorScene{
		tileInCursor: &tiles.TileBuilder{
			Kind: tiles.KindWall,

			Key:     keys[0],
			SourceX: 0,
			SourceY: 0,
		},
		playerPos: playerPos,
		grid:      level.Grid,

		showInfo: true,
	}
}

func (s *EditorScene) InitializedOnce() bool {
	return s.initializedScene
}

func (s *EditorScene) Init() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	s.initializedScene = true
}

func (s *EditorScene) Update(g *game.Game) {
	s.tileInCursor.Key = draw.Keys()[s.keyIndex]
	s.tileInCursor.SourceX = s.sourcePos.X
	s.tileInCursor.SourceY = s.sourcePos.Y

	s.Input()

	ts := constants.TileSize
	x, y := ebiten.CursorPosition()
	x, y = x/ts, y/ts
	if (x >= 0 && x < constants.TilesPerWidth) && (y >= 0 && y < constants.TilesPerHeight) {
		s.cursorPos = units.Vec{X: x, Y: y}
	}
}

func (s *EditorScene) Draw(screen *ebiten.Image) {
	DrawBackground(screen)

	s.DrawTiles(screen)
	s.DrawPlayerPos(screen)

	s.DrawTileCursor(screen)

	draw.DrawCursor(screen)

	if s.showInfo {
		s.PrintInfo(screen)
	}
}

func (s *EditorScene) PrintInfo(screen *ebiten.Image) {
	rotationStr := s.tileInCursor.Transform.String()
	mirroredX, mirroredY := s.tileInCursor.XMirrored, s.tileInCursor.YMirrored

	keyStr := fmt.Sprintf(`"%s" [%v/%v]`, s.tileInCursor.Key, s.keyIndex+1, len(keys))

	w, h := draw.CountTiledImage(draw.GetImage(s.tileInCursor.Key), constants.TileSize)
	maxSource := units.Vec{X: w - 1, Y: h - 1}

	math.Sqrt(2)

	info := fmt.Sprintf(
		`Cursor position: %s

Key: %s
	Source position: %s
          - Max: %s
Rotation: %s
Mirrored: {X: %v, Y: %v}

FPS: %.2f

`,
		s.cursorPos.String(), keyStr, s.sourcePos.String(), maxSource.String(),
		rotationStr, mirroredX, mirroredY, ebiten.CurrentFPS(),
	)

	if !s.saved {
		info += "+"
	} else {
		info += "Saved!"
	}

	ebitenutil.DebugPrint(screen, info)
}
