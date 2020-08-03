// +build js

package util

import (
	"syscall/js"
)

func IsFullscreen() bool {
	fn := "IsFullscreen"
	handler := js.Global().Get(fn)
	if handler.IsUndefined() || handler.IsNull() {
		return false
	}
	val := js.Global().Call(fn)
	if val.IsUndefined() || val.IsNull() {
		return false
	}
	return val.Bool()
}

func SetFullscreen(fullscreen bool) {
	fn := "SetFullscreen"
	handler := js.Global().Get(fn)
	if handler.IsUndefined() || handler.IsNull() {
		return
	}
	js.Global().Call(fn, fullscreen)
}
