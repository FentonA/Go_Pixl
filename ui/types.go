package ui

import (
	"pixl/apptype"
	"pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindo fyne.Window
	State     *apptype.State
	Swatches  []*swatch.Swatch
}
