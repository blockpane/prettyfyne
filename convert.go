package prettyfyne

import (
	"fyne.io/fyne"
	"image/color"
)

// customTheme wraps a PrettyTheme and provides the fyne.Theme interface
type customTheme struct {
	pt *PrettyTheme
}

func (c customTheme) BackgroundColor() color.Color {
	return c.pt.BackgroundColor
}

func (c customTheme) ButtonColor() color.Color {
	return c.pt.ButtonColor
}

func (c customTheme) DisabledButtonColor() color.Color {
	return c.pt.DisabledButtonColor
}

func (c customTheme) HyperlinkColor() color.Color {
	return c.pt.HyperlinkColor
}

func (c customTheme) TextColor() color.Color {
	return c.pt.TextColor
}

func (c customTheme) DisabledTextColor() color.Color {
	return c.pt.DisabledTextColor
}

func (c customTheme) IconColor() color.Color {
	return c.pt.IconColor
}

func (c customTheme) DisabledIconColor() color.Color {
	return c.pt.DisabledIconColor
}

func (c customTheme) PlaceHolderColor() color.Color {
	return c.pt.PlaceHolderColor
}

func (c customTheme) PrimaryColor() color.Color {
	return c.pt.PrimaryColor
}

func (c customTheme) HoverColor() color.Color {
	return c.pt.HoverColor
}

func (c customTheme) FocusColor() color.Color {
	return c.pt.FocusColor
}

func (c customTheme) ScrollBarColor() color.Color {
	return c.pt.ScrollBarColor
}

func (c customTheme) ShadowColor() color.Color {
	return c.pt.ShadowColor
}

func (c customTheme) TextSize() int {
	return c.pt.TextSize
}

func (c customTheme) TextFont() fyne.Resource {
	return c.pt.TextFont
}

func (c customTheme) TextBoldFont() fyne.Resource {
	return c.pt.TextBoldFont
}

func (c customTheme) TextItalicFont() fyne.Resource {
	return c.pt.TextItalicFont
}

func (c customTheme) TextBoldItalicFont() fyne.Resource {
	return c.pt.TextBoldItalicFont
}

func (c customTheme) TextMonospaceFont() fyne.Resource {
	return c.pt.TextMonospaceFont
}

func (c customTheme) Padding() int {
	return c.pt.Padding
}

func (c customTheme) IconInlineSize() int {
	return c.pt.IconInlineSize
}

func (c customTheme) ScrollBarSize() int {
	return c.pt.ScrollBarSize
}

func (c customTheme) ScrollBarSmallSize() int {
	return c.pt.ScrollBarSmallSize
}
