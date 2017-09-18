package processing

import (
	"image/color"
)

type RedTransformation struct{}

func NewRedTransformation() *RedTransformation {
	return &RedTransformation{}
}

func (t *RedTransformation) TransformPixel(colour color.RGBA) color.RGBA {
	return color.RGBA{
		R: colour.R,
		G: 0,
		B: 0,
		A: colour.A,
	}
}
