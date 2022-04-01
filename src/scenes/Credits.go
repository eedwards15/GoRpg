package scenes

import (
	"GoRpg/src/entities"
	"GoRpg/src/models"
	"GoRpg/src/systems"
	"GoRpg/src/ui"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"reflect"
	"time"
)

type Credits struct {
	menuFont      font.Face
	FontSize      int
	currentCredit *entities.AnimatedUiText
	creditList    []*entities.AnimatedUiText
	displaying    []*entities.AnimatedUiText

	lastUpdate      time.Time
	nextIndex       int
	animationTicker int64
}

func NewCredits() *Credits {
	m := &Credits{}
	m.FontSize = 60

	return m
}

func (m *Credits) Init() {
	fontFromSystem := systems.ASSETSYSTEM.Assets["Global"].Fonts["KennySquare"]
	m.menuFont, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
		Size:    float64(m.FontSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	m.nextIndex = 0
	m.animationTicker = 700

	for i := 0; i < len(systems.ASSETSYSTEM.Credit); i++ {
		record := systems.ASSETSYSTEM.Credit[i]

		v := reflect.ValueOf(record)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			s := fmt.Sprintf("%s: %v Testing", typeOfS.Field(i).Name, v.Field(i).Interface())
			start := models.NewVector(float64(ui.CenterTextXPos(s, m.menuFont, systems.WINDOWMANAGER.ScreenWidth)), float64(m.FontSize))
			end := models.NewVector(start.Xpos, float64(systems.WINDOWMANAGER.ScreenHeight+m.FontSize))
			m.creditList = append(m.creditList, entities.NewAnimatedUiText(*end, *start, m.animationTicker, true, m.menuFont, s))
		}

		//Empty Space
		for i := 0; i < 2; i++ {
			s := fmt.Sprintf("             ")
			start := models.NewVector(float64(ui.CenterTextXPos(s, m.menuFont, systems.WINDOWMANAGER.ScreenWidth)), float64(m.FontSize))
			end := models.NewVector(start.Xpos, float64(systems.WINDOWMANAGER.ScreenHeight+m.FontSize))
			m.creditList = append(m.creditList, entities.NewAnimatedUiText(*end, *start, m.animationTicker, true, m.menuFont, s))
		}
	}

}

func (m *Credits) Update() error {
	if time.Now().Sub(m.lastUpdate).Milliseconds() >= m.animationTicker && m.nextIndex < len(m.creditList) {
		m.lastUpdate = time.Now()
		m.displaying = append(m.displaying, m.creditList[m.nextIndex])
		m.nextIndex += 1
	}

	for i := 0; i < len(m.displaying); i++ {
		m.displaying[i].Update()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) || inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		systems.SCENEMANAGER.Pop()
	}
	return nil
}

func (m Credits) Draw(screen *ebiten.Image) {
	for i := 0; i < len(m.displaying); i++ {
		m.displaying[i].Draw(screen)
	}

}

func (m Credits) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}
