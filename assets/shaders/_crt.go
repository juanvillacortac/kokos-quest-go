// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package shaders

var (
	Boost            float32
	GrilleOpacity    float32
	ScanlinesOpacity float32
	VignetteOpacity  float32
	ScanlinesSpeed   float32
	AberrationAmount float32 = 0
	AberrationSpeed  float32 = 1

	ShowGrille     bool = true
	ShowScanlines  bool = true
	ShowVignette   bool = true
	ShowCurvature  bool = true
	MoveAberration bool = false

	ScreenSize vec2
)

func CRTCurveUV(uv vec2) vec2 {
	if ShowCurvature {
		uv = uv*2.0 - 1.0
		offset := abs(uv.yx) / vec2{6.0, 4.0}
		uv = uv + uv*offset*offset
		uv = uv*0.5 + 0.5
	}
	return uv
}

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	dir := normalize(position.xy - Cursor)
	clr := imageSrc0UnsafeAt(texCoord)

	samples := [...]float{
		-22, -14, -8, -4, -2, 2, 4, 8, 14, 22,
	}
	sum := clr
	for i := 0; i < len(samples); i++ {
		pos := texCoord + dir*samples[i]/imageSrcTextureSize()
		sum += imageSrc2At(pos)
	}
	sum /= 10 + 1

	dist := distance(position.xy, Cursor)
	t := clamp(dist/256, 0, 1)
	return mix(clr, sum, t)
}
