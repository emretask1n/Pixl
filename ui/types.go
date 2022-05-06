package ui

import (
	"fyne.io/fyne/v2"
	"github.com/emretask1n/pixl/apptype"
	"github.com/emretask1n/pixl/swatch"
	"github.com/emretask1n/pixl/pxcanvas"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
