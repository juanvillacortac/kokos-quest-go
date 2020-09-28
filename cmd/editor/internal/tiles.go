package editor

import (
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"

	"github.com/thoas/go-funk"
)

func (s *EditorScene) clearGrid() {
	s.grid = make(map[units.Vec]tiles.TileBuilder)
	s.status = UNSAVED
}

func (s *EditorScene) placeTile() {
	if s.tileInCursor != nil {
		s.grid[s.cursorPos] = *s.tileInCursor
	}
	s.status = UNSAVED
}

func (s *EditorScene) removeTile() {
	delete(s.grid, s.cursorPos)
	s.status = UNSAVED
}

func (s *EditorScene) setPlayerPos() {
	if s.grid[s.cursorPos].Key == "" {
		s.playerPos = &units.Vec{X: s.cursorPos.X, Y: s.cursorPos.Y}
		s.status = UNSAVED
	}
}

func (s *EditorScene) incrementKeyIndex() {
	if s.keyIndex != len(keys)-1 {
		s.keyIndex++
	} else {
		s.keyIndex = 0
	}
	s.sourcePos = units.Vec{}
	resetTileTransform(s.tileInCursor)
}

func (s *EditorScene) decrementKeyIndex() {
	if s.keyIndex != 0 {
		s.keyIndex--
	} else {
		s.keyIndex = len(keys) - 1
	}
	s.sourcePos = units.Vec{}
	resetTileTransform(s.tileInCursor)
}

func (s *EditorScene) incrementSourcePos() {
	cx, cy := draw.CountTiledImage(draw.GetImage(s.tileInCursor.Key), units.TileSize)
	if s.sourcePos.X < cx {
		s.sourcePos.X++
		if s.sourcePos.X == cx {
			s.sourcePos.X = 0
			s.sourcePos.Y++
			if s.sourcePos.Y == cy {
				s.sourcePos.Y = 0
			}
		}
	}
	resetTileTransform(s.tileInCursor)
}

func (s *EditorScene) decrementSourcePos() {
	cx, cy := draw.CountTiledImage(draw.GetImage(s.tileInCursor.Key), units.TileSize)
	if s.sourcePos.X >= 0 {
		s.sourcePos.X--
		if s.sourcePos.X < 0 {
			s.sourcePos.X = cx - 1
			s.sourcePos.Y--
			if s.sourcePos.Y < 0 {
				s.sourcePos.Y = cy - 1
			}
		}
	}
	resetTileTransform(s.tileInCursor)
}

func (s *EditorScene) chooseTile() {
	bgTile := s.grid[s.cursorPos]
	index := funk.IndexOf(keys, bgTile.Key)
	if index != -1 {
		s.keyIndex = index
		s.sourcePos = units.Vec{X: bgTile.SourceX, Y: bgTile.SourceY}
		s.tileInCursor = &bgTile
	}
}

func resetTileTransform(tile *tiles.TileBuilder) {
	tile.Transform = tiles.TransformEq
	tile.XMirrored, tile.YMirrored = false, false
}

func rotateTile(tile *tiles.TileBuilder) {
	if tile.Transform != tiles.Transform270 {
		tile.Transform++
	} else {
		tile.Transform = tiles.TransformEq
	}
}

func mirrorXTile(tile *tiles.TileBuilder) {
	tile.XMirrored = !tile.XMirrored
}

func mirrorYTile(tile *tiles.TileBuilder) {
	tile.YMirrored = !tile.YMirrored
}
