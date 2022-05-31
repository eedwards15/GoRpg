package entities

import (
	"GoRpg/src/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpawnPoint struct {
	components.Sprite
	components.Transform
}

func NewSpawnPoint(x, y float64, image *ebiten.Image) *SpawnPoint {
	p := &SpawnPoint{}
	p.Xpos = x
	p.Ypos = y
	p.SetSprite(image)
	return p
}

func (spawnPointClass *SpawnPoint) Draw(image *ebiten.Image) {
	playerDrawOptions := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterLinear,
	}
	playerDrawOptions.GeoM.Translate(spawnPointClass.Xpos, spawnPointClass.Ypos)
	image.DrawImage(spawnPointClass.CurrentImage, playerDrawOptions)
}

func (spawnPointClass *SpawnPoint) Update() {

}
