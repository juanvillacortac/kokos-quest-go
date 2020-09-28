package shaders

var (
	Time       float
	Cursor     vec2
	ScreenSize vec2
)

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	dir := normalize(position.xy - Cursor)
	clr := imageSrc0UnsafeAt(texCoord)

	samples := [...]float{
		-22, -14, -8, -4, -2, 2, 4, 8, 14, 22,
	}
	sum := clr
	for i := 0; i < len(samples); i++ {
		pos := texCoord + dir*samples[i]/ScreenSize
		sum += imageSrc0At(pos)
	}
	sum /= 10 + 1

	dist := distance(position.xy, Cursor)
	t := clamp(dist/256, 0, 1)
	return mix(clr, sum, t)
}
