package ui

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jamesrr39/goutil/image-processing/imageprocessingutil"
	"github.com/jamesrr39/goutil/must"
	"github.com/jamesrr39/gtk3-image-gallery-demo-app/domain"
	"github.com/jamesrr39/image-processing/processing"
)

type ImageCard struct {
	*AppWindow
	*domain.ImageInfo
	imageContainer *gtk.Box
	imageWidget    *gtk.Image
}

func NewImageCard(appWindow *AppWindow, imageInfo *domain.ImageInfo) *ImageCard {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	must.Must(err)
	return &ImageCard{appWindow, imageInfo, box, nil}
}

func (c *ImageCard) Render() gtk.IWidget {

	picture, err := c.AppWindow.ImageDAL.GetImage(c.ImageInfo.RelativePath)
	if nil != err {
		label, err := gtk.LabelNew("couldn't create image. Error: " + err.Error())
		must.Must(err)
		return label
	}

	winWidth, winHeight := c.win.GetSize()

	newSize := imageprocessingutil.GetResizeSize(
		imageprocessingutil.Size{
			Width:  picture.Bounds().Max.X,
			Height: picture.Bounds().Max.Y,
		},
		imageprocessingutil.Size{
			Width:  winWidth - 100,
			Height: winHeight - 30,
		})

	resizedPicture := imaging.Resize(picture, newSize.Width, newSize.Height, imaging.Lanczos)

	c.setImage(resizedPicture)

	// create buttons

	normalButton, err := gtk.ButtonNewWithLabel("Normal")
	must.Must(err)
	normalButton.Connect("clicked", func() {
		c.setImage(resizedPicture)
	})

	simpleButton, err := gtk.ButtonNewWithLabel("Simple")
	must.Must(err)
	simpleButton.Connect("clicked", func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewSimpleColourTransformation(8)))
	})

	negativeButton, err := gtk.ButtonNewWithLabel("Negative")
	must.Must(err)
	negativeButton.Connect("clicked", func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewNegativeTransformation()))
	})

	greyscaleButton, err := gtk.ButtonNewWithLabel("Greyscale")
	must.Must(err)
	greyscaleButton.Connect("clicked", func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewGreyScaleTransformation()))
	})

	redButton, err := gtk.ButtonNewWithLabel("Red")
	must.Must(err)
	redButton.Connect("clicked", func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewRedTransformation()))
	})

	greenButton, err := gtk.ButtonNewWithLabel("Green")
	must.Must(err)
	greenButton.Connect("clicked", func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewGreenTransformation()))
	})

	blueButton, err := gtk.ButtonNewWithLabel("Blue")
	must.Must(err)
	blueButton.Connect("clicked", func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewBlueTransformation()))
	})

	pictureModeButtonsBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	must.Must(err)

	pictureModeButtonsBox.PackStart(normalButton, false, false, 0)
	pictureModeButtonsBox.PackStart(simpleButton, false, false, 0)
	pictureModeButtonsBox.PackStart(negativeButton, false, false, 0)
	pictureModeButtonsBox.PackStart(greyscaleButton, false, false, 0)
	pictureModeButtonsBox.PackStart(redButton, false, false, 0)
	pictureModeButtonsBox.PackStart(greenButton, false, false, 0)
	pictureModeButtonsBox.PackStart(blueButton, false, false, 0)

	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	must.Must(err)
	hbox.PackStart(c.imageContainer, false, false, 0)
	hbox.PackStart(pictureModeButtonsBox, false, false, 0)

	return hbox
}

func (c *ImageCard) setImage(resizedPicture image.Image) {
	pixbuf, err := PixBufFromImage(resizedPicture)
	must.Must(err)

	imageWidget, err := gtk.ImageNewFromPixbuf(pixbuf)
	must.Must(err)
	if nil != c.imageWidget {
		c.imageWidget.Destroy()
	}
	c.imageWidget = imageWidget
	c.imageContainer.PackStart(c.imageWidget, false, false, 0)
	c.imageContainer.ShowAll()
}
