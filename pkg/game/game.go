package game

import (
	"fmt"
	"sync"

	"kokos_quest/pkg/audio"
	"kokos_quest/pkg/draw"
	"kokos_quest/pkg/global"
	"kokos_quest/pkg/util"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

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

	scene Scene

	debuggable bool
}

func NewGame(title string, width, height, scale int, debuggable bool) (*Game, error) {
	global.ScreenWidth = width * scale
	global.ScreenHeight = height * scale

	ebiten.SetWindowSize(global.ScreenWidth, global.ScreenHeight)
	ebiten.SetWindowTitle(title)

	global.Debug = debuggable

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
		if audioErr := audio.Finalize(); audioErr != nil {
			err = audioErr
		}
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
	return g.width, g.height
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

	g.initScene()

	g.scene.Update(g)

	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		util.SetFullscreen(!util.IsFullscreen())
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		if g.debuggable {
			global.Debug = !global.Debug
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.resourceLoadedCh != nil {
		ebitenutil.DebugPrint(screen, "Loading resources...")
		return
	}
	if g.scene == nil {
		return
	}

	g.scene.Draw(screen)

	if global.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()))
	}
}
