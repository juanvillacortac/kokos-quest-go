package tiles

import "kokos_quest/pkg/draw"

func EmptyTile() *Tile {
	return &Tile{
		kind: KindEmpty,
	}
}

func NewWall() *Tile {
	img := draw.GetTiledImage("Wall", 1, 0)
	return &Tile{
		kind:      KindWall,
		tileImage: img,
	}
}
