package world

import (
	"kokos_quest/pkg/constants"
	"kokos_quest/pkg/player"
	"kokos_quest/pkg/tiles"
	"kokos_quest/pkg/units"
)

type WorldBuilder struct {
	mapDims   units.Dims
	grid      map[units.Vec]tiles.TileBuilder
	playerPos units.Vec
}

func NewBuilder(dims units.Dims) WorldBuilder {
	return WorldBuilder{
		mapDims: dims,
		grid:    make(map[units.Vec]tiles.TileBuilder, 0),
	}
}

func (b *WorldBuilder) SetGrid(grid map[units.Vec]tiles.TileBuilder) {
	b.grid = grid
}

func (b *WorldBuilder) SetPlayer(pos units.Vec) {
	b.playerPos = pos
}

func (b *WorldBuilder) Build() *World {
	grid := make(map[units.Vec]tiles.Tile)
	for pos, tb := range b.grid {
		grid[pos] = tb.Build()
	}

	playerDir := units.Right
	if b.playerPos.X > constants.TilesPerWidth/2 {
		playerDir = units.Left
	}

	player := player.NewPlayer(b.playerPos, playerDir)

	space := initSpace(b.playerPos, grid)
	space.Add(player.Rect())

	return &World{
		mapDims: b.mapDims,
		grid:    grid,
		player:  player,
		space:   space,
	}
}
