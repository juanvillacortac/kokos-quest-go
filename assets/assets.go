package assets

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thoas/go-funk"
)

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -nocompress -pkg=assets sprites sounds music levels shaders
//go:generate gofmt -s -w .

// Load file keys given in a directory
func LoadKeys(dir string, suffixes []string) []string {
	keys, _ := AssetDir(dir)
	return funk.FilterString(keys, func(str string) bool {
		for _, suffix := range suffixes {
			if strings.HasSuffix(str, suffix) {
				return true
			}
		}
		return false
	})
}

// Load file keys given in a directory, stripping the extension suffix
func LoadKeysStripped(dir string, suffixes []string) []string {
	keys := funk.Map(LoadKeys(dir, suffixes), func(str string) string {
		for _, suffix := range suffixes {
			str = strings.TrimSuffix(str, suffix)
		}
		return str
	}).([]string)
	sort.Strings(keys)
	return keys
}

// Return a sprite in a slice of bytes
func Image(key string) []byte {
	return MustAsset(fmt.Sprintf("sprites/%s.png", key))
}

// Return a shader in a slice of bytes
func Shader(key string) []byte {
	return MustAsset(fmt.Sprintf("shaders/_%s.go", key))
}

// Return a level string
func Level(key string) []byte {
	return MustAsset(fmt.Sprintf("levels/%s", key))
}
