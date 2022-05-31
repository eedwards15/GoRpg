// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package systems

import (
	"GoRpg/src/components"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
	"math"
)

type Camera struct {
	ViewPort        f64.Vec2
	Position        f64.Vec2
	ZoomFactor      int
	Rotation        int
	FollowTransform *components.Transform
}

func (c *Camera) Follow(vectorToFollow *components.Transform) {

	c.FollowTransform = vectorToFollow
}
func (c *Camera) Update() {
	if c.FollowTransform != nil {
		newX := c.FollowTransform.Xpos - c.viewportCenter()[0]
		newY := c.FollowTransform.Ypos - (c.viewportCenter()[1]) - 20

		if newX <= -(c.viewportCenter()[0]) {
			newX = -(c.viewportCenter()[0])
		}

		if newX >= c.viewportCenter()[0] {
			newX = (c.viewportCenter()[0])
		}

		if newY <= -(c.viewportCenter()[1]) {
			newY = -(c.viewportCenter()[1])
		}

		if newY >= c.ViewPort[1]+c.viewportCenter()[1]+32 {
			newY = c.ViewPort[1] + c.viewportCenter()[1] + 32
		}

		c.Position[0] = newX
		c.Position[1] = newY
	}
}

func (c *Camera) String() string {
	return fmt.Sprintf(
		"T: %.1f, R: %d, S: %d",
		c.Position, c.Rotation, c.ZoomFactor,
	)
}

func (c *Camera) viewportCenter() f64.Vec2 {
	return f64.Vec2{
		c.ViewPort[0] * 0.5,
		c.ViewPort[1] * 0.5,
	}
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	// We want to scale and rotate around center of image / screen
	m.Translate(-c.viewportCenter()[0], -c.viewportCenter()[1])
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.viewportCenter()[0], c.viewportCenter()[1])
	return m
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(float64(posX), float64(posY))
	} else {
		// When scaling it can happend that matrix is not invertable
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
	c.Rotation = 0
	c.ZoomFactor = 0
}
