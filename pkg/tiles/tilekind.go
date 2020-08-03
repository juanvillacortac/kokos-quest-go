package tiles

type TileKind int

const (
	KindEmpty TileKind = iota
	KindWall
)

func (k TileKind) String() string {
	switch k {
	case KindEmpty:
		return "KindEmpty"
	case KindWall:
		return "KindWall"
	default:
		return "-- unknown --"
	}
}
