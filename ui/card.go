package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type Card interface {
	Render() gtk.IWidget
}
