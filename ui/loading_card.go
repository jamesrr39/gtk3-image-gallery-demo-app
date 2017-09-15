package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/jamesrr39/goutil/must"
)

type LoadingCard struct {
}

func NewLoadingCard() *LoadingCard {
	return &LoadingCard{}
}

func (c *LoadingCard) Render() gtk.IWidget {
	label, err := gtk.LabelNew("Loading")
	must.Must(err)

	return label
}
