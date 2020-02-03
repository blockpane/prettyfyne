package prettyfyne

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"gopkg.in/yaml.v2"
	"image/color"
)

type PrettyTheme struct {
	BackgroundColor        color.Color
	ButtonColor            color.Color
	DisabledButtonColor    color.Color
	HyperlinkColor         color.Color
	TextColor              color.Color
	DisabledTextColor      color.Color
	IconColor              color.Color
	DisabledIconColor      color.Color
	PlaceHolderColor       color.Color
	PrimaryColor           color.Color
	HoverColor             color.Color
	FocusColor             color.Color
	ScrollBarColor         color.Color
	ShadowColor            color.Color
	TextSize               int
	TextFont               fyne.Resource
	textFontPath           string
	TextBoldFont           fyne.Resource
	textBoldFontPath       string
	TextItalicFont         fyne.Resource
	textItalicFontPath     string
	TextBoldItalicFont     fyne.Resource
	textBoldItalicFontPath string
	TextMonospaceFont      fyne.Resource
	textMonospaceFontPath  string
	Padding                int
	IconInlineSize         int
	ScrollBarSize          int
	ScrollBarSmallSize     int
}

// DefaultTheme returns the default fyne dark theme, which provides an easy starting point for customization
func DefaultTheme() *PrettyTheme {
	return &PrettyTheme{
		BackgroundColor:     theme.DarkTheme().BackgroundColor(),
		ButtonColor:         theme.DarkTheme().ButtonColor(),
		DisabledButtonColor: theme.DarkTheme().DisabledButtonColor(),
		HyperlinkColor:      theme.DarkTheme().HyperlinkColor(),
		TextColor:           theme.DarkTheme().TextColor(),
		DisabledTextColor:   theme.DarkTheme().DisabledIconColor(),
		IconColor:           theme.DarkTheme().IconColor(),
		DisabledIconColor:   theme.DarkTheme().DisabledIconColor(),
		PlaceHolderColor:    theme.DarkTheme().PlaceHolderColor(),
		PrimaryColor:        theme.DarkTheme().PrimaryColor(),
		HoverColor:          theme.DarkTheme().HoverColor(),
		FocusColor:          theme.DarkTheme().FocusColor(),
		ScrollBarColor:      theme.DarkTheme().ScrollBarColor(),
		ShadowColor:         theme.DarkTheme().ShadowColor(),
		TextSize:            theme.DarkTheme().TextSize(),
		TextFont:            theme.DarkTheme().TextFont(),
		TextBoldFont:        theme.DarkTheme().TextBoldFont(),
		TextItalicFont:      theme.DarkTheme().TextItalicFont(),
		TextBoldItalicFont:  theme.DarkTheme().TextBoldItalicFont(),
		TextMonospaceFont:   theme.DarkTheme().TextMonospaceFont(),
		Padding:             theme.DarkTheme().Padding(),
		IconInlineSize:      theme.DarkTheme().IconInlineSize(),
		ScrollBarSize:       theme.DarkTheme().ScrollBarSize(),
		ScrollBarSmallSize:  theme.DarkTheme().ScrollBarSmallSize(),
	}
}

// DefaultLightTheme returns the default fyne light theme, which provides an easy starting point for customization
func DefaultLightTheme() *PrettyTheme {
	return &PrettyTheme{
		BackgroundColor:     theme.LightTheme().BackgroundColor(),
		ButtonColor:         theme.LightTheme().ButtonColor(),
		DisabledButtonColor: theme.LightTheme().DisabledButtonColor(),
		HyperlinkColor:      theme.LightTheme().HyperlinkColor(),
		TextColor:           theme.LightTheme().TextColor(),
		DisabledTextColor:   theme.LightTheme().DisabledIconColor(),
		IconColor:           theme.LightTheme().IconColor(),
		DisabledIconColor:   theme.LightTheme().DisabledIconColor(),
		PlaceHolderColor:    theme.LightTheme().PlaceHolderColor(),
		PrimaryColor:        theme.LightTheme().PrimaryColor(),
		HoverColor:          theme.LightTheme().HoverColor(),
		FocusColor:          theme.LightTheme().FocusColor(),
		ScrollBarColor:      theme.LightTheme().ScrollBarColor(),
		ShadowColor:         theme.LightTheme().ShadowColor(),
		TextSize:            theme.LightTheme().TextSize(),
		TextFont:            theme.LightTheme().TextFont(),
		TextBoldFont:        theme.LightTheme().TextBoldFont(),
		TextItalicFont:      theme.LightTheme().TextItalicFont(),
		TextBoldItalicFont:  theme.LightTheme().TextBoldItalicFont(),
		TextMonospaceFont:   theme.LightTheme().TextMonospaceFont(),
		Padding:             theme.LightTheme().Padding(),
		IconInlineSize:      theme.LightTheme().IconInlineSize(),
		ScrollBarSize:       theme.LightTheme().ScrollBarSize(),
		ScrollBarSmallSize:  theme.LightTheme().ScrollBarSmallSize(),
	}
}


// ToFyneTheme converts a PrettyTheme to the fyne.Theme interface that can be applied to the app
func (pt PrettyTheme) ToFyneTheme() fyne.Theme {
	return customTheme{&pt}
}

func (pt PrettyTheme) MarshalYaml() ([]byte, error) {
	return yaml.Marshal(pt.ToConfig())
}

func (pt PrettyTheme) ToConfig() PrettyThemeConfig {
	r := func(c color.Color) *color.RGBA {
		r, g, b, a := c.RGBA()
		return &color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	}

	return PrettyThemeConfig{
		BackgroundColor:        r(pt.BackgroundColor),
		ButtonColor:            r(pt.ButtonColor),
		DisabledButtonColor:    r(pt.DisabledButtonColor),
		HyperlinkColor:         r(pt.HyperlinkColor),
		TextColor:              r(pt.TextColor),
		DisabledTextColor:      r(pt.DisabledTextColor),
		IconColor:              r(pt.IconColor),
		DisabledIconColor:      r(pt.DisabledIconColor),
		PlaceHolderColor:       r(pt.PlaceHolderColor),
		PrimaryColor:           r(pt.PrimaryColor),
		HoverColor:             r(pt.HoverColor),
		FocusColor:             r(pt.FocusColor),
		ScrollBarColor:         r(pt.ScrollBarColor),
		ShadowColor:            r(pt.ShadowColor),
		TextSize:               pt.TextSize,
		TextFontPath:           pt.textFontPath,
		TextFont:               pt.TextFont.Name(),
		TextBoldFontPath:       pt.textBoldFontPath,
		TextBoldFont:           pt.TextBoldFont.Name(),
		TextItalicFontPath:     pt.textItalicFontPath,
		TextItalicFont:         pt.TextItalicFont.Name(),
		TextBoldItalicFontPath: pt.textBoldItalicFontPath,
		TextBoldItalicFont:     pt.TextBoldItalicFont.Name(),
		TextMonospaceFontPath:  pt.textMonospaceFontPath,
		TextMonospaceFont:      pt.TextMonospaceFont.Name(),
		Padding:                pt.Padding,
		IconInlineSize:         pt.IconInlineSize,
		ScrollBarSize:          pt.ScrollBarSize,
		ScrollBarSmallSize:     pt.ScrollBarSmallSize,
	}
}

var ExampleMaterialLight = PrettyTheme{
	BackgroundColor:        color.RGBA{255,255,255,255},
	ButtonColor:            color.RGBA{0,0,0,10},
	DisabledButtonColor:    color.RGBA{0,0,0,30},
	HyperlinkColor:         color.RGBA{100,181,246,253},
	TextColor:              color.RGBA{0,0,0,222},
	DisabledTextColor:      color.RGBA{94,94,94,255},
	IconColor:              color.RGBA{100,181,246,253},
	DisabledIconColor:      color.RGBA{0,0,0,97},
	PlaceHolderColor:       color.RGBA{178,178,178,255},
	PrimaryColor:           color.RGBA{166,212,250,255},
	HoverColor:             color.RGBA{0,0,0,30},
	FocusColor:             color.RGBA{100,181,246,255},
	ScrollBarColor:         color.RGBA{0,0,0,153},
	ShadowColor:            color.RGBA{0,0,0,66},
	TextSize:               13,
	TextFont:               theme.LightTheme().TextFont(),
	TextBoldFont:           theme.LightTheme().TextBoldFont(),
	TextItalicFont:         theme.LightTheme().TextItalicFont(),
	TextBoldItalicFont:     theme.LightTheme().TextBoldItalicFont(),
	TextMonospaceFont:      theme.LightTheme().TextMonospaceFont(),
	Padding:                4,
	IconInlineSize:         24,
	ScrollBarSize:          12,
	ScrollBarSmallSize:     3,
}

var ExampleEveryDayIsHalloween = PrettyTheme{
	BackgroundColor:        color.RGBA{R: 30, G: 30, B: 30, A: 255},
	ButtonColor:            color.RGBA{R: 20, G: 20, B: 20, A: 255},
	DisabledButtonColor:    color.RGBA{R: 15, G: 15, B: 17, A: 255},
	HyperlinkColor:         color.RGBA{R:0xaa, G:0x64, B:0x14, A:0x40},
	TextColor:              color.RGBA{R: 200, G: 200, B: 200, A: 255},
	DisabledTextColor:      color.RGBA{R: 155, G: 155, B: 155, A: 127},
	IconColor:              color.RGBA{R: 150, G: 80, B: 0, A: 255},
	DisabledIconColor:      color.RGBA{R: 155, G: 155, B: 155, A: 127},
	PlaceHolderColor:       color.RGBA{R: 150, G: 80, B: 0, A: 255},
	PrimaryColor:           color.RGBA{R: 110, G: 40, B: 0, A: 127},
	HoverColor:             color.RGBA{R: 0, G: 0, B: 0, A: 255},
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