package processing

import (
	"image/color"
)

type GreyScaleTransformation struct {
}

// NewGreyScaleTransformation creates a transformation.
func NewGreyScaleTransformation() *GreyScaleTransformation {
	return &GreyScaleTransformation{}
}

func (t *GreyScaleTransformation) TransformPixel(pixelColour color.RGBA) color.RGBA {
	n := (float64(pixelColour.R) + float64(pixelColour.G) + float64(pixelColour.B)) / 3

	return color.RGBA{
		R: uint8(n),
		G: uint8(n),
		B: uint8(n),
		A: pixelColour.A,
	}
}
