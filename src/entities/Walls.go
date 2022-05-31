package entities

import "GoRpg/src/components"

type Wall struct {
	components.Transform
	Width  float64
	Height float64
}

func NewWall(x, y, w, h float64) *Wall {
	p := &Wall{}
	p.Xpos = x
	p.Ypos = y
	p.Width = w
	p.Height = h
	return p
}
