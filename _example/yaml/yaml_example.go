package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/blockpane/prettyfyne"
	"os"
	"time"
)

// override only a few from the defaults:
var y = `
background_color:
  r: 42
  g: 42
  b: 42
  a: 255
shadow_color:
  r: 42
  g: 42
  b: 42
  a: 255
focus_color:
  r: 62
  g: 62
  b: 62
  a: 255
hover_color:
  r: 0
  g: 0
  b: 0
  a: 255
padding: 8
icon_inline_size: 32
`

func main() {
	a := app.New()
	w := a.NewWindow("Pretty Fyne")

	pt, _, _ := prettyfyne.UnmarshalYaml([]byte(y))
	yamlEntry := widget.NewMultiLineEntry()
	yamlEntry.SetText(y)

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Theme From Yaml Example"),
		widget.NewButtonWithIcon("Quit", theme.CancelIcon(), func() {
			a.Quit()
		}),
		fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(200, 500)),
			widget.NewScrollContainer(
				yamlEntry,
			)),
	))

	go func(){
		// delay for app to be visible:
		time.Sleep(100 * time.Millisecond)
		fyne.CurrentApp().Settings().SetTheme(pt.ToFyneTheme())
		w.Content().Refresh()
	}()

	w.SetFixedSize(true)
	w.ShowAndRun()
}
