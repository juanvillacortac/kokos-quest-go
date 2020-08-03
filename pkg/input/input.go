package input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Button int

const (
	UP Button = iota
	DOWN
	LEFT
	RIGHT
	A
	B
	START
)

func (btn Button) String() string {
	switch btn {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	case A:
		return "A"
	case B:
		return "B"
	case START:
		return "START"
	default:
		return "UNKNOWN"
	}
}

type GamepadAxis struct {
	Id    int
	Value float64
}

func (ga *GamepadAxis) IsTrue(gpId int) bool {
	return ga.Value == ebiten.GamepadAxis(gpId, ga.Id)
}

var (
	KeysMap = map[Button]ebiten.Key{
		UP:    ebiten.KeyUp,
		DOWN:  ebiten.KeyDown,
		LEFT:  ebiten.KeyLeft,
		RIGHT: ebiten.KeyRight,
		A:     ebiten.KeyZ,
		B:     ebiten.KeyX,
		START: ebiten.KeyEnter,
	}

	GamepadAxisMap = map[Button]*GamepadAxis{
		UP: {
			Id:    1,
			Value: -1.0,
		},
		DOWN: {
			Id:    1,
			Value: 1.0,
		},
		LEFT: {
			Id:    0,
			Value: -1.0,
		},
		RIGHT: {
			Id:    0,
			Value: 1.0,
		},
	}

	GamepadButtonsMap = map[Button]ebiten.GamepadButton{
		A:     ebiten.GamepadButton0,
		B:     ebiten.GamepadButton1,
		START: ebiten.GamepadButton2,
	}
)

func GetGamepadId() (int, bool) {
	gamepads := ebiten.GamepadIDs()
	if len(gamepads) > 0 {
		return gamepads[0], true
	}
	return 0, false
}

func IsKeyPressed(btn Button) bool {
	gpId, _ := GetGamepadId()
	mapping, ok := GamepadButtonsMap[btn]
	var gp bool
	if ok {
		gp = ebiten.IsGamepadButtonPressed(gpId, mapping)
	} else {
		gp = GamepadAxisMap[btn].IsTrue(gpId)
	}
	return ebiten.IsKeyPressed(KeysMap[btn]) || gp
}

func IsKeyJustPressed(btn Button) bool {
	gpId, _ := GetGamepadId()
	mapping, ok := GamepadButtonsMap[btn]
	var gp bool
	if ok {
		gp = inpututil.IsGamepadButtonJustPressed(gpId, mapping)
	}
	return inpututil.IsKeyJustPressed(KeysMap[btn]) || gp
}
