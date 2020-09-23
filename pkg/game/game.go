package game

import (
	"fmt"
	"os"
	"sync"

	"kokos_quest/pkg/audio"
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/global"
	"kokos_quest/pkg/units"
	"kokos_quest/pkg/util"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var shader = []byte(`
package main

var Time float
var Cursor vec2
var ScreenSize vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	center := ScreenSize / 2
	amount := (center - Cursor) / 10 / ScreenSize
	var clr vec4
	clr.r = imageSrc0At(texCoord + amount).r
	clr.g = imageSrc0UnsafeAt(texCoord).g
	clr.b = imageSrc0At(texCoord - amount).b
	return clr
}
`)

type Scene interface {
	InitializedOnce() bool

	Init()
	Update(g *Game)
	Draw(screen *ebiten.Image)
}

type Game struct {
	width, height int

	initializedResources sync.Once
	resourceLoadedCh     chan error

	initialSceneCallback func() Scene

	scene  Scene
	shader *ebiten.Shader

	time int

	debuggable bool
}

func NewGame(title string, width, height, scale int, debuggable bool) (*Game, error) {
	global.ScreenSize = units.Dims{
		W: width * scale,
		H: height * scale,
	}

	ebiten.SetWindowSize(global.ScreenSize.W, global.ScreenSize.H)
	ebiten.SetWindowTitle(title)

	global.Config.Debug = debuggable

	game := &Game{
		width:            width,
		height:           height,
		debuggable:       debuggable,
		resourceLoadedCh: make(chan error),
	}

	if err := game.InitResources(); err != nil {
		return nil, err
	}

	game.InitStuff()

	return game, nil
}

func (g *Game) SetInitialScene(callback func() Scene) {
	g.initialSceneCallback = callback
}

func (g *Game) InitResources() error {
	var err error
	g.initializedResources.Do(func() {
		go func() {
			defer close(g.resourceLoadedCh)
			if err := draw.LoadImages(); err != nil {
				g.resourceLoadedCh <- err
				return
			}
			if err := audio.Load(); err != nil {
				g.resourceLoadedCh <- err
				return
			}
		}()
		s, shaderErr := ebiten.NewShader(shader)
		if shaderErr != nil {
			err = shaderErr
		}
		if audioErr := audio.Finalize(); audioErr != nil {
			err = audioErr
		}
		g.shader = s
	})
	return err
}

func (g *Game) initScene() {
	if g.scene == nil {
		if g.initialSceneCallback != nil {
			g.scene = g.initialSceneCallback()
		} else {
			g.scene = NewTitleScene()
		}
		return
	}

	if !g.scene.InitializedOnce() {
		g.scene.Init()
	}
}

func (g *Game) SetScene(scene Scene) {
	g.scene = scene
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	if global.Config.HighDpi {
		return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
	} else {
		return g.width, g.height
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.resourceLoadedCh != nil {
		select {
		case err := <-g.resourceLoadedCh:
			if err != nil {
				return err
			}
			g.resourceLoadedCh = nil
		default:
		}
	}

	g.time++

	g.initScene()

	g.scene.Update(g)

	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		util.SetFullscreen(!util.IsFullscreen())
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		if g.debuggable {
			global.Config.Debug = !global.Config.Debug
		}
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		global.Config.Shader = !global.Config.Shader
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		global.Config.HighDpi = !global.Config.HighDpi
	case inpututil.IsKeyJustPressed(ebiten.KeyEscape):
		os.Exit(0)
	}

	return nil
}

func getScaleFactor(screen, rect []int) float64 {
	screenAspect := float64(screen[0]) / float64(screen[1])
	rectAspect := float64(rect[0]) / float64(rect[1])

	var scaleFactor float64
	if screenAspect > rectAspect {
		scaleFactor = float64(screen[1]) / float64(rect[1])
	} else {
		scaleFactor = float64(screen[0]) / float64(rect[0])
	}

	return scaleFactor
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.resourceLoadedCh != nil {
		ebitenutil.DebugPrint(screen, "Loading resources...")
		return
	}
	if g.scene == nil {
		return
	}

	sw, sh := screen.Size()
	global.ScreenSize = units.Dims{
		W: sw,
		H: sh,
	}
	global.VirtualScreenSize = units.Dims{
		W: g.width,
		H: g.height,
	}

	framebuffer, _ := ebiten.NewImage(g.width, g.height, ebiten.FilterDefault)

	g.scene.Draw(framebuffer)

	op := &ebiten.DrawImageOptions{}

	global.ScaleFactor = getScaleFactor([]int{sw, sh}, []int{g.width, g.height})

	if global.Config.HighDpi {
		// Move the images's center to the upper left corner.
		op.GeoM.Translate(float64(-g.width)/2, float64(-g.height)/2)

		// The image is just too big. Adjust the scale.
		op.GeoM.Scale(global.ScaleFactor, global.ScaleFactor)

		// Scale the image by the device ratio so that the rendering result can be same
		// on various (different-DPI) environments.
		scale := ebiten.DeviceScaleFactor()
		op.GeoM.Scale(scale, scale)

		// Move the image's center to the screen's center.
		op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	}

	if global.Config.Debug {
		ebitenutil.DebugPrint(framebuffer, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))
	}

	if global.Config.Shader {
		opShaded := &ebiten.DrawRectShaderOptions{GeoM: op.GeoM}
		opShaded.Uniforms = map[string]interface{}{
			"Time": float32(g.time) / 60,
			"Cursor": []float32{
				float32(global.CursorPosition().X),
				float32(global.CursorPosition().Y),
			},
			"ScreenSize": []float32{
				float32(global.VirtualScreenSize.W),
				float32(global.VirtualScreenSize.H),
			},
		}
		opShaded.Images[0] = framebuffer
		screen.DrawRectShader(g.width, g.height, g.shader, opShaded)
	} else {
		screen.DrawImage(framebuffer, op)
	}
}
