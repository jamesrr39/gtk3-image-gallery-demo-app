package processing

import (
	"image/color"
)

type GreenTransformation struct{}

func NewGreenTransformation() *GreenTransformation {
	return &GreenTransformation{}
}

func (t *GreenTransformation) TransformPixel(colour color.RGBA) color.RGBA {
	return color.RGBA{
		R: 0,
		G: colour.G,
		B: 0,
		A: colour.A,
	}
}
