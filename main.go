package main

import (
	"GoRpg/src"
	"GoRpg/src/scenes"
	"GoRpg/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	systems.InitAssetSystem()
	systems.InitWindowManger(1920, 1080)
	systems.InitSceneManager()
	systems.InitMusicSystem(systems.ASSETSYSTEM.Assets["Title"].BackgroundMusic)

	g := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Go RPG")

	systems.SCENEMANAGER.Push(scenes.NewMainMenu())

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
