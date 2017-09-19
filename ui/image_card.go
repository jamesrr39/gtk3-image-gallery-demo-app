package ui

import (
	"image"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	gotk3extra "github.com/jamesrr39/go-gtk-extra/gotk3-extra"
	"github.com/jamesrr39/goutil/debounce"
	"github.com/jamesrr39/goutil/image-processing/imageprocessingutil"
	"github.com/jamesrr39/goutil/must"
	"github.com/jamesrr39/gtk3-image-gallery-demo-app/domain"
	"github.com/jamesrr39/image-processing/processing"
)

type ImageCard struct {
	*AppWindow
	*domain.ImageInfo
	imageContainer        *gtk.Box
	imageWidget           *gtk.Image
	currentTransformation func()
}

func NewImageCard(appWindow *AppWindow, imageInfo *domain.ImageInfo) *ImageCard {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	must.Must(err)
	return &ImageCard{appWindow, imageInfo, box, nil, nil}
}

func (c *ImageCard) Render() gtk.IWidget {

	picture, err := c.AppWindow.ImageDAL.GetImage(c.ImageInfo.RelativePath)
	if nil != err {
		label, err := gtk.LabelNew("couldn't create image. Error: " + err.Error())
		must.Must(err)
		return label
	}

	resizedPicture := c.resizePicture(picture)

	// create render functions
	renderNormal := func() {
		c.setImage(resizedPicture)
	}

	renderSimple := func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewSimpleColourTransformation(8)))
	}

	renderNegative := func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewNegativeTransformation()))
	}

	renderGreyScale := func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewGreyScaleTransformation()))
	}

	renderRed := func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewRedTransformation()))
	}

	renderGreen := func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewGreenTransformation()))
	}

	renderBlue := func() {
		c.setImage(processing.Transform(resizedPicture, processing.NewBlueTransformation()))
	}

	// create buttons

	normalButton, err := gtk.ButtonNewWithLabel("Normal")
	must.Must(err)
	normalButton.Connect("clicked", func() {
		renderNormal()
		c.currentTransformation = renderNormal
	})

	simpleButton, err := gtk.ButtonNewWithLabel("Simple")
	must.Must(err)
	simpleButton.Connect("clicked", func() {
		renderSimple()
		c.currentTransformation = renderSimple
	})

	negativeButton, err := gtk.ButtonNewWithLabel("Negative")
	must.Must(err)
	negativeButton.Connect("clicked", func() {
		renderNegative()
		c.currentTransformation = renderNegative
	})

	greyscaleButton, err := gtk.ButtonNewWithLabel("Greyscale")
	must.Must(err)
	greyscaleButton.Connect("clicked", func() {
		renderGreyScale()
		c.currentTransformation = renderGreyScale
	})

	redButton, err := gtk.ButtonNewWithLabel("Red")
	must.Must(err)
	redButton.Connect("clicked", func() {
		renderRed()
		c.currentTransformation = renderRed
	})

	greenButton, err := gtk.ButtonNewWithLabel("Green")
	must.Must(err)
	greenButton.Connect("clicked", func() {
		renderGreen()
		c.currentTransformation = renderGreen
	})

	blueButton, err := gtk.ButtonNewWithLabel("Blue")
	must.Must(err)
	blueButton.Connect("clicked", func() {
		renderBlue()
		c.currentTransformation = renderBlue
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

	c.currentTransformation = renderNormal
	c.setImage(resizedPicture)

	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	must.Must(err)
	hbox.PackStart(c.imageContainer, false, false, 0)
	hbox.PackStart(pictureModeButtonsBox, false, false, 0)

	var previousAllocation *gtk.Allocation
	debouncer := debounce.NewDebouncer(time.Millisecond * 100)

	c.AppWindow.win.Connect("size-allocate", func() {
		debouncer.Run(func() {
			allocation := c.imageContainer.GetAllocation()
			if previousAllocation == nil || allocation.GetWidth() != previousAllocation.GetWidth() || allocation.GetHeight() != previousAllocation.GetHeight() {
				previousAllocation = allocation
				resizedPicture = c.resizePicture(picture)
				glib.IdleAdd(func() {
					c.currentTransformation()
					c.imageContainer.ShowAll()
				})
			}
		})
	})

	return hbox
}

func (c *ImageCard) resizePicture(picture image.Image) image.Image {
	winWidth, winHeight := c.win.GetSize()

	newSize := imageprocessingutil.GetResizeSize(
		imageprocessingutil.Size{
			Width:  picture.Bounds().Max.X,
			Height: picture.Bounds().Max.Y,
		},
		imageprocessingutil.Size{
			Width:  winWidth - 100,
			Height: winHeight - 100,
		})

	resizedPicture := imaging.Resize(picture, newSize.Width, newSize.Height, imaging.Lanczos)

	return resizedPicture
}

func (c *ImageCard) setImage(resizedPicture image.Image) {
	pixbuf, err := gotk3extra.PixBufFromImage(resizedPicture)
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
