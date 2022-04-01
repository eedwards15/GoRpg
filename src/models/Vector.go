package models

type Vector struct {
	Xpos float64
	Ypos float64
}

func NewVector(xpos, ypos float64) *Vector {
	v := &Vector{
		Xpos: xpos,
		Ypos: ypos,
	}

	return v
}
