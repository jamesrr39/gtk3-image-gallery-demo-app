package processing

import (
	"image"
	"image/color"
)

type Transformation interface {
	TransformPixel(color.RGBA) color.RGBA
}

func Transform(picture image.Image, t Transformation) image.Image {

	newPicture := image.NewRGBA(picture.Bounds())
	for y := 0; y < picture.Bounds().Max.Y; y++ {
		for x := 0; x < picture.Bounds().Max.X; x++ {
			pixelColour := picture.At(x, y)
			r, g, b, a := pixelColour.RGBA()

			ratio8Bit32Bit := float64(255) / float64(65336)

			eightBitColour := color.RGBA{
				R: uint8(float64(r) * ratio8Bit32Bit),
				G: uint8(float64(g) * ratio8Bit32Bit),
				B: uint8(float64(b) * ratio8Bit32Bit),
				A: uint8(float64(a) * ratio8Bit32Bit),
			}

			newPicture.SetRGBA(x, y, t.TransformPixel(eightBitColour))
		}
	}
	return newPicture
}
