package src

import (
	"GoRpg/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame() *Game {
	g := &Game{}
	return g
}

func (gameClass *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}

func (gameClass *Game) Update() error {
	if systems.SCENEMANAGER.CurrentScene != nil {
		systems.SCENEMANAGER.CurrentScene.Update()
	}
	return nil
}

func (gameClass *Game) Draw(screen *ebiten.Image) {
	if systems.SCENEMANAGER.CurrentScene != nil {
		systems.SCENEMANAGER.CurrentScene.Draw(screen)
	}
}
