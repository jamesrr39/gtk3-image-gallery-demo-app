package processing

import (
	"image/color"
	"math"
)

type SimpleColourTransformation struct {
	shadesPerColourFactor float64
}

// NewSimpleColourTransformation creates a transformation.
// `numberOfShadesPerColour` is the number of shades of each colour are allowed in the transformation.
// example: `4` would mean all the shades of red are rounded into 4 shades (0-63, 64-127, 128-191, 192-255)
func NewSimpleColourTransformation(numberOfShadesPerColour uint8) *SimpleColourTransformation {
	scale := float64(256) / float64(numberOfShadesPerColour)

	return &SimpleColourTransformation{scale}
}

func (t *SimpleColourTransformation) TransformPixel(pixelColour color.RGBA) color.RGBA {

	return color.RGBA{
		R: t.transformColour(pixelColour.R),
		G: t.transformColour(pixelColour.G),
		B: t.transformColour(pixelColour.B),
		A: pixelColour.A,
	}
}

func (t *SimpleColourTransformation) transformColour(colourValue uint8) uint8 {
	value := math.Floor(float64(colourValue)/float64(t.shadesPerColourFactor)) * float64(t.shadesPerColourFactor)
	if value > 255 {
		return 255
	}
	return uint8(value)
}
