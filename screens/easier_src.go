package screens

import (
	"fmt"

	"github.com/blockpane/prettyfyne"
)

func ToEasierGoSource(config prettyfyne.PrettyThemeConfig) string {
	return fmt.Sprintf(`package main

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

        myTheme :=  prettyfyne.PrettyTheme{
                BackgroundColor:        %#v,
                ButtonColor:            %#v,
                DisabledButtonColor:    %#v,
                HyperlinkColor:         %#v,
                TextColor:              %#v,
                DisabledTextColor:      %#v,
                IconColor:              %#v,
                DisabledIconColor:      %#v,
                PlaceHolderColor:       %#v,
                PrimaryColor:           %#v,
                HoverColor:             %#v,
                FocusColor:             %#v,
                ScrollBarColor:         %#v,
                ShadowColor:            %#v,
                TextSize:               %v,
                TextFont:               theme.DarkTheme().TextFont(),
                TextBoldFont:           theme.DarkTheme().TextBoldFont(),
                TextItalicFont:         theme.DarkTheme().TextItalicFont(),
                TextBoldItalicFont:     theme.DarkTheme().TextBoldItalicFont(),
                TextMonospaceFont:      theme.DarkTheme().TextMonospaceFont(),
                Padding:                %v,
                IconInlineSize:         %v,
                ScrollBarSize:          %v,
                ScrollBarSmallSize:     %v,
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
`,
		config.BackgroundColor,
		config.ButtonColor,
		config.DisabledButtonColor,
		config.HyperlinkColor,
		config.TextColor,
		config.DisabledTextColor,
		config.IconColor,
		config.DisabledIconColor,
		config.PlaceHolderColor,
		config.PrimaryColor,
		config.HoverColor,
		config.FocusColor,
		config.ScrollBarColor,
		config.ShadowColor,
		config.TextSize,
		config.Padding,
		config.IconInlineSize,
		config.ScrollBarSize,
		config.ScrollBarSmallSize,
	)
}
