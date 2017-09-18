package processing

import (
	"image/color"
)

type BlueTransformation struct{}

func NewBlueTransformation() *BlueTransformation {
	return &BlueTransformation{}
}

func (t *BlueTransformation) TransformPixel(colour color.RGBA) color.RGBA {
	return color.RGBA{
		R: 0,
		G: 0,
		B: colour.B,
		A: colour.A,
	}
}
