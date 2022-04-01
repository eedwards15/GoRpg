package entities

import (
	"GoRpg/src/components"
	"GoRpg/src/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type AnimatedUiText struct {
	Starting    components.Transform
	Ending      components.Transform
	amountToAdd *models.Vector
	Current     *models.Vector
	//Create an option obj that can be passed in.
	milliseconds  int64
	ShouldFadeOut bool
	font          font.Face
	text          string
	IsComplete    bool
}

func NewAnimatedUiText(starting models.Vector, ending models.Vector, milliseconds int64, shouldFadeOut bool, face font.Face, text string) *AnimatedUiText {
	a := &AnimatedUiText{}
	a.Current = models.NewVector(starting.Xpos, starting.Ypos)

	a.Starting = components.Transform{
		Vector: starting,
	}

	a.Ending = components.Transform{
		Vector: ending,
	}

	a.milliseconds = milliseconds
	a.ShouldFadeOut = shouldFadeOut
	a.font = face
	a.text = text
	x := (a.Ending.Xpos - a.Starting.Xpos) / float64(a.milliseconds)
	y := (a.Ending.Ypos - a.Starting.Ypos) / float64(a.milliseconds)
	a.amountToAdd = models.NewVector(x, y)
	a.IsComplete = false
	return a
}

func (a AnimatedUiText) Draw(screen *ebiten.Image) {
	text.Draw(screen, a.text, a.font, int(a.Current.Xpos), int(a.Current.Ypos), color.White)
}

//Not working Correctly.
func (a *AnimatedUiText) Update() error {
	if a.Current.Xpos != a.Ending.Xpos || ((a.Current.Ypos <= a.Ending.Ypos || a.Current.Ypos >= a.Ending.Ypos) && a.Current.Ypos > -60) {
		a.Current.Xpos += a.amountToAdd.Xpos
		a.Current.Ypos += a.amountToAdd.Ypos
	} else {
		a.IsComplete = true
	}
	return nil
}
