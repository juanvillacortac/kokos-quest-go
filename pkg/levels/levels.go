package levels

import (
	"fmt"

	"kokos_quest/assets"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"
	"kokos_quest/pkg/world"

	"github.com/kelindar/binary"
)

func parse(file LevelFile) *world.WorldBuilder {
	world := world.NewBuilder(units.Dims{W: 20, H: 20})
	world.SetGrid(file.Grid)
	world.SetPlayer(file.PlayerPos)
	return &world
}

func Load(level_num uint) *world.World {
	level := newLevelFile(assets.Level(fmt.Sprintf("level_%v", level_num)))
	return parse(level).Build()
}

func LoadLiteral(level_key string) *world.World {
	level := newLevelFile(assets.Level(level_key))
	return parse(level).Build()
}

type LevelFile struct {
	Grid      map[units.Vec]tiles.TileBuilder
	PlayerPos units.Vec
}

func newLevelFile(data []byte) LevelFile {
	var l LevelFile
	if err := binary.Unmarshal(data, &l); err != nil {
		panic("Error parsing level buffer")
	}
	return l
}
