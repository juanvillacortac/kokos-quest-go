package shaders

var (
	Time       float
	Cursor     vec2
	ScreenSize vec2
)

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	center := ScreenSize / 2
	amount := (center - Cursor) / 10 / ScreenSize
	var clr vec4
	clr.r = imageSrc0At(texCoord + amount).r
	clr.g = imageSrc0UnsafeAt(texCoord).g
	clr.b = imageSrc0At(texCoord - amount).b
	return clr
}
