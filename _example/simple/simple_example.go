package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	pf "github.com/blockpane/prettyfyne"
	"golang.org/x/image/colornames"
	"image/color"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Pretty Fyne - Simple Example")

	// get a copy of the default fyne theme:
	prettyTheme := pf.DefaultTheme()

	// Change the font
	label := widget.NewLabel("Using custom font Arial.ttf from system path")
	var err error
	prettyTheme.TextFont, err = pf.LoadFont("", "Arial.ttf", pf.FontDefault)
	if err != nil {
		label.SetText("Couldn't load custom font, using the default\n"+err.Error())
	}

	// Customize the sizing
	prettyTheme.TextSize = 18
	prettyTheme.Padding = 5
	prettyTheme.IconInlineSize = 24

	// Change a few colors
	prettyTheme.TextColor = colornames.Whitesmoke
	prettyTheme.DisabledTextColor = colornames.Dimgray
	prettyTheme.BackgroundColor = color.RGBA{ R:20, G: 20, B: 20, A: 128}
	prettyTheme.ButtonColor = colornames.Black
	prettyTheme.HoverColor = colornames.Olivedrab
	prettyTheme.IconColor = colornames.Chocolate
	prettyTheme.ShadowColor = color.RGBA{ R:10, G: 10, B: 10, A: 255}

	yamlString, _ := prettyTheme.MarshalYaml()
	yamlEntry := widget.NewMultiLineEntry()
	yamlEntry.SetText(string(yamlString))

	disabledButton := widget.NewButtonWithIcon("Disabled", theme.InfoIcon(), func(){})
	disabledButton.Disable()
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Custom Theme!"),
		widget.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Quit", theme.CancelIcon(), func() {
				a.Quit()
			}),
			disabledButton,
			layout.NewSpacer(),
		),
		label,
		fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(500, 400)),
			widget.NewScrollContainer(
				yamlEntry,
		)),
	))

	go func(){
		w.Hide()
		time.Sleep(100 * time.Millisecond)
		// call ToFyneTheme to return the interface:
		fyne.CurrentApp().Settings().SetTheme(prettyTheme.ToFyneTheme())
		w.Content().Refresh()
		w.Show()
	}()

	w.SetFixedSize(true)
	w.ShowAndRun()
}
