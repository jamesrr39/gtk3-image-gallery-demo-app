package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/jamesrr39/goutil/must"
	"github.com/jamesrr39/gtk3-app/dal"
	"github.com/jamesrr39/gtk3-app/ui"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	rootDir := kingpin.Arg(
		"root directory",
		"the folder you want to show images from").Required().String()
	kingpin.Parse()

	imageDAL := dal.NewImageDAL(*rootDir)
	err := imageDAL.ScanRootPath()
	must.Must(err)

	gtk.Init(nil)

	ui.NewAppWindow(imageDAL)

	gtk.Main()
}
