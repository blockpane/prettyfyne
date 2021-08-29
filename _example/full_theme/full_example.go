package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/blockpane/prettyfyne"
	"image/color"
	"time"
)

func main()  {
	a := app.New()
	w := a.NewWindow("Pretty Fyne - PrettyTheme Example")

	myTheme := prettyfyne.PrettyTheme{
		BackgroundColor:        color.RGBA{R: 30, G: 30, B: 30, A: 255},
		ButtonColor:            color.RGBA{R: 20, G: 20, B: 20, A: 255},
		DisabledButtonColor:    color.RGBA{R: 15, G: 15, B: 17, A: 255},
		HyperlinkColor:         color.RGBA{R: 170, G:100, B:20, A: 64},
		TextColor:              color.RGBA{R: 200, G: 200, B: 200, A: 255},
		DisabledTextColor:      color.RGBA{R: 155, G: 155, B: 155, A: 127},
		IconColor:              color.RGBA{R: 150, G: 80, B: 0, A: 255},
		DisabledIconColor:      color.RGBA{R: 155, G: 155, B: 155, A: 127},
		PlaceHolderColor:       color.RGBA{R: 150, G: 80, B: 0, A: 255},
		PrimaryColor:           color.RGBA{R: 110, G: 40, B: 0, A: 127},
		HoverColor:             color.RGBA{R: 110, G: 40, B: 0, A: 127},
		FocusColor:             color.RGBA{R: 99, G: 99, B: 99, A: 255},
		ScrollBarColor:         color.RGBA{R: 35, G: 35, B: 35, A: 8},
		ShadowColor:            color.RGBA{R: 0, G: 0, B: 0, A: 64},
		TextSize:               13,
		TextFont:               theme.DarkTheme().TextFont(),
		TextBoldFont:           theme.DarkTheme().TextBoldFont(),
		TextItalicFont:         theme.DarkTheme().TextItalicFont(),
		TextBoldItalicFont:     theme.DarkTheme().TextBoldItalicFont(),
		TextMonospaceFont:      theme.DarkTheme().TextMonospaceFont(),
		Padding:                4,
		IconInlineSize:         24,
		ScrollBarSize:          10,
		ScrollBarSmallSize:     4,
	}

	hello := widget.NewLabel("Hello Pretty Fyne!")

	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	// applying a theme before the window is running will cause a panic, aggressively wait for it to be present
	go func(){
		for {
			time.Sleep(50 * time.Millisecond)
			if hello.Visible() {
				break
			}
		}
		fyne.CurrentApp().Settings().SetTheme(myTheme.ToFyneTheme())
	}()

	w.CenterOnScreen()
	w.ShowAndRun()
}
