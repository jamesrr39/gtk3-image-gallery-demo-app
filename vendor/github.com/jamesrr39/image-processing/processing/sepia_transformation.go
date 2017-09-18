package processing

import (
	"image/color"
)

var (
	//maxRedFraction   = float64(112) / 255
	maxRedFraction = float64(100) / 255
	//maxGreenFraction = float64(66) / 255
	maxGreenFraction = float64(50) / 255
	maxBlueFraction  = float64(30) / 255
)

type SepiaTransformation struct {
}

// NewGreyScaleTransformation creates a transformation.
func NewSepiaTransformation() *SepiaTransformation {
	return &SepiaTransformation{}
}

func (t *SepiaTransformation) TransformPixel(pixelColour color.RGBA) color.RGBA {
	return color.RGBA{
		R: uint8(float64(pixelColour.R) * maxRedFraction),
		G: uint8(float64(pixelColour.G) * maxGreenFraction),
		B: uint8(float64(pixelColour.B) * maxBlueFraction),
		A: pixelColour.A,
	}
}
