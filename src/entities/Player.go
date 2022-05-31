package entities

import (
	"GoRpg/src/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	components.Transform
	components.Sprite
	IsColliding bool
}

func InitPlayer(x, y float64, image *ebiten.Image) *Player {
	p := &Player{}
	p.Xpos = x
	p.Ypos = y
	p.SetSprite(image)
	p.IsColliding = false
	return p
}

func (player *Player) Draw(image *ebiten.Image) {
	playerDrawOptions := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterLinear,
	}
	playerDrawOptions.GeoM.Translate(player.Xpos, player.Ypos)
	image.DrawImage(player.CurrentImage, playerDrawOptions)
}

func (player *Player) Update() {

	println(player.IsColliding)

	if ebiten.IsKeyPressed(ebiten.KeyD) && !inpututil.IsKeyJustReleased(ebiten.KeyD) {
		player.Xpos += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) && !inpututil.IsKeyJustReleased(ebiten.KeyS) {
		player.Ypos += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && !inpututil.IsKeyJustReleased(ebiten.KeyW) {
		player.Ypos -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && !inpututil.IsKeyJustReleased(ebiten.KeyA) {
		player.Xpos -= 5
	}
}
