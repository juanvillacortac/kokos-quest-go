package editor

import (
	"fmt"
	"io/ioutil"
	"os"

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

func loadLevel(path string) (levels.LevelFile, error) {
	l := levels.LevelFile{
		Grid:      make(map[units.Vec]tiles.TileBuilder, 0),
		PlayerPos: units.Vec{X: -1, Y: -1},
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return l, nil
	}
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return l, fmt.Errorf("[ERROR] %s", err)
	}
	if err = binary.Unmarshal(buff, &l); err != nil {
		return l, fmt.Errorf("[ERROR] %s", err)
	}
	return l, nil
}

func (s *EditorScene) saveLevel(path string) {
	if s.saved {
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
	if err := ioutil.WriteFile(path, buff, 0644); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	audio.PlaySE("Jump")
	s.saved = true
}
