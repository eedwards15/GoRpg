package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type IScene interface {
	Init()
	Draw(screen *ebiten.Image)
	Update() error
}
