package ui

import (
	"fmt"

	"github.com/disintegration/imaging"
	"github.com/gotk3/gotk3/gtk"
	"github.com/jamesrr39/goutil/must"
)

type GalleryListingCard struct {
	*AppWindow
}

func NewGalleryListingCard(appWindow *AppWindow) *GalleryListingCard {
	return &GalleryListingCard{appWindow}
}

func (c *GalleryListingCard) Render() gtk.IWidget {
	imageInfos := c.AppWindow.ImageDAL.GetAll()
	if 0 == len(imageInfos) {
		label, err := gtk.LabelNew("No images found")
		must.Must(err)

		return label
	}

	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	must.Must(err)

	for _, imageInfo := range imageInfos {
		label, err := gtk.LabelNew(imageInfo.RelativePath)
		must.Must(err)

		var imageWidget gtk.IWidget
		picture, err := c.GetImage(imageInfo.RelativePath)
		if nil != err {
			imageWidget, err = gtk.LabelNew(fmt.Sprintf("couldn't get '%s'. Error: %s",
				imageInfo.RelativePath,
				err))
			must.Must(err)
		} else {
			resizedPicture := imaging.Resize(picture, 200, 300, imaging.Lanczos)
			pixBuf, err := PixBufFromImage(resizedPicture)
			imageWidget, err = gtk.LabelNew(fmt.Sprintf("couldn't create an image for '%s'. Error: %s",
				imageInfo.RelativePath,
				err))
			must.Must(err)

			imageWidget, err = gtk.ImageNewFromPixbuf(pixBuf)
			must.Must(err)
		}

		vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
		must.Must(err)
		vbox.PackStart(imageWidget, false, false, 0)
		vbox.PackStart(label, false, false, 0)

		hbox.PackStart(vbox, false, false, 5)
	}

	/*
		scrollWin, err := gtk.ScrolledWindowNew(nil, nil)
		must.Must(err)
		scrollWin.SetPolicy(gtk.POLICY_ALWAYS, gtk.POLICY_ALWAYS)

		scrollWin.Add(hbox)
	*/
	return hbox
}
