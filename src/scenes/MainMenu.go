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
	"os"
)

type MainMenu struct {
	menuFont         font.Face
	loaded           bool
	fontSize         int
	offset           int
	currentMenuIndex int
}

func NewMainMenu() *MainMenu {
	m := &MainMenu{}
	m.loaded = false
	return m
}

func (m *MainMenu) Init() {
	if m.loaded {
		systems.MUSICSYSTEM.LoadSong(systems.ASSETSYSTEM.Assets["Title"].BackgroundMusic).PlaySong()
	}

	if !m.loaded {
		systems.MUSICSYSTEM.PlaySong()

		println("Init Main Menu")
		fontFromSystem := systems.ASSETSYSTEM.Assets["Global"].Fonts["KennySquare"]
		m.fontSize = 60
		m.offset = 120
		m.currentMenuIndex = 0
		m.menuFont, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
			Size:    float64(m.fontSize),
			DPI:     72,
			Hinting: font.HintingFull,
		})
		m.loaded = true
	}

}

func (m *MainMenu) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		if m.currentMenuIndex == 1 {
			systems.SCENEMANAGER.Push(NewCredits())
		}

		if m.currentMenuIndex == 0 {
			systems.SCENEMANAGER.Push(NewMainScene())
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		if m.currentMenuIndex == 0 {
			m.currentMenuIndex += 1
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		if m.currentMenuIndex == 1 {
			m.currentMenuIndex -= 1
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(1)
	}

	return nil
}

func (m MainMenu) Draw(screen *ebiten.Image) {

	text.Draw(screen, "Go Rpg", m.menuFont, ui.CenterTextXPos("Go Rpg", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), m.fontSize, color.White)

	if m.currentMenuIndex == 0 {
		text.Draw(screen, "Start", m.menuFont, ui.CenterTextXPos("Start", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), (m.fontSize*2)+m.offset, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	} else {
		text.Draw(screen, "Start", m.menuFont, ui.CenterTextXPos("Start", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), (m.fontSize*2)+m.offset, color.White)
	}

	if m.currentMenuIndex == 1 {
		text.Draw(screen, "Credits", m.menuFont, ui.CenterTextXPos("Credits", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), (m.fontSize*4)+m.offset, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	} else {
		text.Draw(screen, "Credits", m.menuFont, ui.CenterTextXPos("Credits", m.menuFont, systems.WINDOWMANAGER.ScreenWidth), (m.fontSize*4)+m.offset, color.White)
	}
}

func (m MainMenu) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}
