package prettyfyne

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// customTheme wraps a PrettyTheme and provides the fyne.Theme interface
type customTheme struct {
	pt *PrettyTheme
}

func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return c.pt.BackgroundColor
	case theme.ColorNameButton:
		return c.pt.ButtonColor
	case theme.ColorNameDisabledButton:
		return c.pt.DisabledButtonColor
	case theme.ColorNameFocus:
		return c.pt.FocusColor
	case theme.ColorNameHover:
		return c.pt.HoverColor
	case theme.ColorNamePlaceHolder:
		return c.pt.PlaceHolderColor
	case theme.ColorNamePrimary:
		return c.pt.PrimaryColor
	case theme.ColorNameScrollBar:
		return c.pt.ScrollBarColor
	case theme.ColorNameShadow:
		return c.pt.ShadowColor
	default:
		return c.pt.PrimaryColor
	}
}

func (c customTheme) Icon(icon fyne.ThemeIconName) fyne.Resource {
	return c.pt.Icon
}

func (c customTheme) Size(size fyne.ThemeSizeName) float32 {
	return c.pt.TextSize
}

func (c customTheme) Font(style fyne.TextStyle) fyne.Resource {
	switch style {
	case fyne.TextStyle{Italic: true}:
		return c.pt.TextItalicFont
	case fyne.TextStyle{Bold: true}:
		return c.pt.TextBoldFont
	case fyne.TextStyle{Bold: true, Italic: true}:
		return c.pt.TextBoldItalicFont
	default:
		return c.pt.TextFont
	}
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

func (c customTheme) TextSize() float32 {
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

func (c customTheme) Padding() float32 {
	return c.pt.Padding
}

func (c customTheme) IconInlineSize() float32 {
	return c.pt.IconInlineSize
}

func (c customTheme) ScrollBarSize() float32 {
	return c.pt.ScrollBarSize
}

func (c customTheme) ScrollBarSmallSize() float32 {
	return c.pt.ScrollBarSmallSize
}
