package screens

import (
	"fmt"

	"github.com/blockpane/prettyfyne"
)

func ToGoSource(config prettyfyne.PrettyThemeConfig) string {
	return fmt.Sprintf(`package customtheme

import (
         "image/color"

         "fyne.io/fyne/v2"
         "fyne.io/fyne/v2/theme"
)

// customTheme is based upon the custom theme example in the fyne_demo application. It is generated from the 
// current template in the prettyfyne theme editor and should be useable with a fyne applicaton without any
// additional requirements.
// It can be applied by calling: "fyne.CurrentApp().Settings().SetTheme(newCustomTheme())" after the app is running.
type customTheme struct {
}

func (customTheme) BackgroundColor() color.Color {
         return %#v
}

func (customTheme) ButtonColor() color.Color {
         return %#v
}

func (customTheme) DisabledButtonColor() color.Color {
         return %#v
}

func (customTheme) HyperlinkColor() color.Color {
         return %#v
}

func (customTheme) TextColor() color.Color {
         return %#v
}

func (customTheme) DisabledTextColor() color.Color {
         return %#v
}

func (customTheme) IconColor() color.Color {
         return %#v
}

func (customTheme) DisabledIconColor() color.Color {
         return %#v
}

func (customTheme) PlaceHolderColor() color.Color {
         return %#v
}

func (customTheme) PrimaryColor() color.Color {
         return %#v
}

func (customTheme) HoverColor() color.Color {
         return %#v
}

func (customTheme) FocusColor() color.Color {
         return %#v
}

func (customTheme) ScrollBarColor() color.Color {
         return %#v
}

func (customTheme) ShadowColor() color.Color {
         return %#v
}

func (customTheme) TextSize() float32 {
         return %v
}

// TODO: for now, demo still returns default fonts
func (customTheme) TextFont() fyne.Resource {
         return theme.DefaultTextFont()
}

func (customTheme) TextBoldFont() fyne.Resource {
         return theme.DefaultTextBoldFont()
}

func (customTheme) TextItalicFont() fyne.Resource {
         return theme.DefaultTextItalicFont()
}

func (customTheme) TextBoldItalicFont() fyne.Resource {
         return theme.DefaultTextBoldItalicFont()
}

func (customTheme) TextMonospaceFont() fyne.Resource {
         return theme.DefaultTextMonospaceFont()
}

func (customTheme) Padding() float32 {
         return %v
}

func (customTheme) IconInlineSize() float32 {
         return %v
}

func (customTheme) ScrollBarSize() float32 {
         return %v
}

func (customTheme) ScrollBarSmallSize() float32 {
         return %v
}

func newCustomTheme() fyne.Theme {
         return &customTheme{}
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
