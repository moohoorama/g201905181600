package mobile

import (
	"github.com/hajimehoshi/ebiten/mobile"

	"github.com/moohoorama/g201905181600/ino"
)

const (
	ScreenWidth  = ino.ScreenWidth
	ScreenHeight = ino.ScreenHeight
)

var (
	running bool
)

func IsRunning() bool {
	return running
}

func Start(scale float64) error {
	running = true
	game, err := ino.NewGame()
	if err != nil {
		return err
	}
	if err := mobile.Start(game.Loop, ino.ScreenWidth, ino.ScreenHeight, scale, ino.Title); err != nil {
		return err
	}
	return nil
}

func Update() error {
	return mobile.Update()
}

func UpdateTouchesOnAndroid(action int, id int, x, y int) {
	mobile.UpdateTouchesOnAndroid(action, id, x, y)
}

func UpdateTouchesOnIOS(phase int, ptr int64, x, y int) {
	mobile.UpdateTouchesOnIOS(phase, ptr, x, y)
}
