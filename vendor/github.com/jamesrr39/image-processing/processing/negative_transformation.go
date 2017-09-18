package processing

import (
	"image/color"
)

type NegativeTransform struct{}

func NewNegativeTransformation() *NegativeTransform {
	return &NegativeTransform{}
}

func (t *NegativeTransform) TransformPixel(colour color.RGBA) color.RGBA {
	return color.RGBA{
		R: 255 - colour.R,
		G: 255 - colour.G,
		B: 255 - colour.B,
		A: colour.A,
	}
}
