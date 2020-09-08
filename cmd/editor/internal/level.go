package editor

import (
	"fmt"

	"kokos_quest/pkg/audio"
	"kokos_quest/pkg/levels"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"

	"github.com/kelindar/binary"
)

// Indios
// Te quiero y nada mas
// Conociendo a rusia
// Doja Cat - Say so
// L.A Espinetta - Ya no mires atras

type LoadLevelFunc func() ([]byte, error)
type SaveLevelFunc func([]byte) error

type SaveStatus int

const (
	UNSAVED SaveStatus = iota
	SAVING
	SAVED
)

func loadLevel(handler LoadLevelFunc) (levels.LevelFile, error) {
	l := levels.LevelFile{
		Grid:      make(map[units.Vec]tiles.TileBuilder, 0),
		PlayerPos: units.Vec{X: -1, Y: -1},
	}

	buff, err := handler()
	if err != nil {
		return l, fmt.Errorf("[ERROR] %s", err)
	}
	if err = binary.Unmarshal(buff, &l); err != nil {
		return l, fmt.Errorf("[ERROR] %s", err)
	}
	return l, nil
}

func (s *EditorScene) saveLevel() {
	if s.status == SAVED {
		return
	}
	if s.playerPos == nil {
		fmt.Println("[ERROR] The world needs a player")
		return
	}
	l := levels.LevelFile{
		Grid:      s.grid,
		PlayerPos: *s.playerPos,
	}
	buff, err := binary.Marshal(l)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	s.status = SAVING

	if err := s.SaveLevelHandler(buff); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	audio.PlaySE("Jump")
	s.status = SAVED
}
