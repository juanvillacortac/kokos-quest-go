package draw

import (
	"bytes"
	"fmt"
	"image"

	"kokos_quest/assets"
	"kokos_quest/pkg/constants"

	"github.com/hajimehoshi/ebiten"
)

const (
	fileSuffix string = ".png"
)

var (
	images     = map[string]*ebiten.Image{}
	spriteKeys []string
)

func init() {
	spriteKeys = assets.LoadKeysStripped("sprites", []string{fileSuffix})
}

// Return keys of sprites loaded in memory
func Keys() []string {
	return spriteKeys
}

// Decode an inlined image at runtime
func DecodeImage(key string) (*ebiten.Image, error) {
	b := assets.Image(key)
	origImg, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	img, _ := ebiten.NewImageFromImage(origImg, ebiten.FilterDefault)
	return img, nil
}

// Load all inlined images to memory
func LoadImages() error {
	for _, f := range spriteKeys {
		img, err := DecodeImage(f)
		if err != nil {
			return err
		}
		images[f] = img
	}
	return nil
}

// Retrieve a image from memory
func GetImage(key string) *ebiten.Image {
	img := images[key]
	if img == nil {
		err := fmt.Errorf(`Image "%s%s" doesn't found on memory`, key, fileSuffix)
		panic(err)
	}
	return img
}

// Retrieve a subimage from memory
func GetSubImage(key string, sx, sy, tile_size int) *ebiten.Image {
	img := GetImage(key)
	dx, dy := sx*tile_size, sy*tile_size
	return img.SubImage(image.Rect(dx, dy, dx+tile_size, dy+tile_size)).(*ebiten.Image)
}

// Shorthand to `GetSubImage` with the global tile size
func GetTiledImage(key string, sx, sy int) *ebiten.Image {
	return GetSubImage(key, sx, sy, constants.TileSize)
}

// Return the count of tiles in an image
func CountTiledImage(img *ebiten.Image, tile_size int) (w int, h int) {
	w, h = img.Size()
	w, h = w/tile_size, h/tile_size
	return
}
