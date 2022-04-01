package scenes

import (
	"GoRpg/src/systems"
	"GoRpg/src/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

type MainScene struct {
	menuFont font.Face
}

func NewMainScene() *MainScene {
	m := &MainScene{}
	return m
}

func (m *MainScene) Init() {
	fontFromSystem := systems.ASSETSYSTEM.Assets["Global"].Fonts["KennySquare"]
	m.menuFont, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})

}

func (m *MainScene) Update() error {

	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) || inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		systems.SCENEMANAGER.Pop()
	}
	return nil
}

func (m MainScene) Draw(screen *ebiten.Image) {
	text.Draw(screen, "MainScene", m.menuFont, ui.CenterTextXPos("MainScene", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), systems.WINDOWMANAGER.YCenter, color.White)

}

func (m MainScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}
