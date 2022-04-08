package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	CurrentImage *ebiten.Image
	ImageWidth   int
	ImageHeight  int
}

func (sprite *Sprite) SetSprite(image *ebiten.Image) {
	sprite.CurrentImage = image
	sprite.ImageWidth, sprite.ImageHeight = sprite.CurrentImage.Size()
}
