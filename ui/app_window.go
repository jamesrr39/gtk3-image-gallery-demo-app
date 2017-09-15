package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/jamesrr39/goutil/must"
	"github.com/jamesrr39/gtk3-app/dal"
)

type AppWindow struct {
	*dal.ImageDAL
	win              *gtk.Window
	outerContainer   *gtk.Box
	contentContainer *gtk.Box
}

func NewAppWindow(dal *dal.ImageDAL) *AppWindow {

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	must.Must(err)

	win.SetTitle("Image Gallery")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	outerContainer, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	must.Must(err)

	win.Add(outerContainer)

	appWindow := &AppWindow{dal, win, outerContainer, nil}

	win.SetDefaultSize(800, 600)

	appWindow.RenderCard(NewGalleryListingCard(appWindow))

	return appWindow
}

func (w *AppWindow) RenderCard(card Card) {

	if nil != w.contentContainer {
		w.contentContainer.Destroy()
	}

	var err error
	w.contentContainer, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	must.Must(err)

	w.contentContainer.Add(card.Render())
	w.outerContainer.Add(w.contentContainer)

	w.win.ShowAll()
}
