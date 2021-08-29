package prettyfyne

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"gopkg.in/yaml.v2"
)

type PrettyTheme struct {
	Icon                   fyne.Resource
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
	TextSize               float32
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
	Padding                float32
	IconInlineSize         float32
	ScrollBarSize          float32
	ScrollBarSmallSize     float32
}

// DefaultTheme returns the default fyne dark theme, which provides an easy starting point for customization
func DefaultTheme() *PrettyTheme {
	return &PrettyTheme{
		Icon:                theme.ComputerIcon(),
		BackgroundColor:     theme.DarkTheme().Color(theme.ColorNameBackground, theme.VariantDark),
		ButtonColor:         theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		DisabledButtonColor: theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		HyperlinkColor:      theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		TextColor:           theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		DisabledTextColor:   theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		IconColor:           theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		DisabledIconColor:   theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		PlaceHolderColor:    theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		PrimaryColor:        theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		HoverColor:          theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		FocusColor:          theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		ScrollBarColor:      theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		ShadowColor:         theme.DarkTheme().Color(theme.ColorNameButton, theme.VariantDark),
		TextSize:            theme.TextSize(),
		TextFont:            theme.DarkTheme().Font(fyne.TextStyle{}),
		TextBoldFont:        theme.DarkTheme().Font(fyne.TextStyle{Bold: true}),
		TextItalicFont:      theme.DarkTheme().Font(fyne.TextStyle{Italic: true}),
		TextBoldItalicFont:  theme.DarkTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
		TextMonospaceFont:   theme.DarkTheme().Font(fyne.TextStyle{Monospace: true}),
		Padding:             theme.Padding(),
		IconInlineSize:      theme.IconInlineSize(),
		ScrollBarSize:       theme.ScrollBarSize(),
		ScrollBarSmallSize:  theme.ScrollBarSmallSize(),
	}
}

// DefaultLightTheme returns the default fyne light theme, which provides an easy starting point for customization
func DefaultLightTheme() *PrettyTheme {
	return &PrettyTheme{
		Icon:                theme.ComputerIcon(),
		BackgroundColor:     theme.LightTheme().Color(theme.ColorNameBackground, theme.VariantLight),
		ButtonColor:         theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		DisabledButtonColor: theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		HyperlinkColor:      theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		TextColor:           theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		DisabledTextColor:   theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		IconColor:           theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		DisabledIconColor:   theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		PlaceHolderColor:    theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		PrimaryColor:        theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		HoverColor:          theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		FocusColor:          theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		ScrollBarColor:      theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		ShadowColor:         theme.LightTheme().Color(theme.ColorNameButton, theme.VariantLight),
		TextSize:            theme.TextSize(),
		TextFont:            theme.LightTheme().Font(fyne.TextStyle{}),
		TextBoldFont:        theme.LightTheme().Font(fyne.TextStyle{Bold: true}),
		TextItalicFont:      theme.LightTheme().Font(fyne.TextStyle{Italic: true}),
		TextBoldItalicFont:  theme.LightTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
		TextMonospaceFont:   theme.LightTheme().Font(fyne.TextStyle{Monospace: true}),
		Padding:             theme.Padding(),
		IconInlineSize:      theme.IconInlineSize(),
		ScrollBarSize:       theme.ScrollBarSize(),
		ScrollBarSmallSize:  theme.ScrollBarSmallSize(),
	}
}

// ToFyneTheme converts a PrettyTheme to the fyne.Theme interface that can be applied to the app
func (pt PrettyTheme) ToFyneTheme() fyne.Theme {
	return customTheme{pt: &pt}
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
	BackgroundColor:     color.RGBA{255, 255, 255, 255},
	ButtonColor:         color.RGBA{0, 0, 0, 10},
	DisabledButtonColor: color.RGBA{0, 0, 0, 30},
	HyperlinkColor:      color.RGBA{100, 181, 246, 253},
	TextColor:           color.RGBA{0, 0, 0, 222},
	DisabledTextColor:   color.RGBA{94, 94, 94, 255},
	IconColor:           color.RGBA{100, 181, 246, 253},
	DisabledIconColor:   color.RGBA{0, 0, 0, 97},
	PlaceHolderColor:    color.RGBA{178, 178, 178, 255},
	PrimaryColor:        color.RGBA{166, 212, 250, 255},
	HoverColor:          color.RGBA{0, 0, 0, 30},
	FocusColor:          color.RGBA{198, 198, 198, 224},
	ScrollBarColor:      color.RGBA{0, 0, 0, 153},
	ShadowColor:         color.RGBA{0, 0, 0, 66},
	TextSize:            13,
	TextFont:            theme.LightTheme().Font(fyne.TextStyle{}),
	TextBoldFont:        theme.LightTheme().Font(fyne.TextStyle{Bold: true}),
	TextItalicFont:      theme.LightTheme().Font(fyne.TextStyle{Italic: true}),
	TextBoldItalicFont:  theme.LightTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
	TextMonospaceFont:   theme.LightTheme().Font(fyne.TextStyle{Monospace: true}),
	Padding:             4,
	IconInlineSize:      24,
	ScrollBarSize:       12,
	ScrollBarSmallSize:  3,
}

var ExampleEveryDayIsHalloween = PrettyTheme{
	BackgroundColor:     color.RGBA{R: 30, G: 30, B: 30, A: 255},
	ButtonColor:         color.RGBA{R: 20, G: 20, B: 20, A: 255},
	DisabledButtonColor: color.RGBA{R: 15, G: 15, B: 17, A: 255},
	HyperlinkColor:      color.RGBA{R: 0xaa, G: 0x64, B: 0x14, A: 0x40},
	TextColor:           color.RGBA{R: 200, G: 200, B: 200, A: 255},
	DisabledTextColor:   color.RGBA{R: 155, G: 155, B: 155, A: 127},
	IconColor:           color.RGBA{R: 150, G: 80, B: 0, A: 255},
	DisabledIconColor:   color.RGBA{R: 155, G: 155, B: 155, A: 127},
	PlaceHolderColor:    color.RGBA{R: 150, G: 80, B: 0, A: 255},
	PrimaryColor:        color.RGBA{R: 110, G: 40, B: 0, A: 127},
	HoverColor:          color.RGBA{R: 0, G: 0, B: 0, A: 255},
	FocusColor:          color.RGBA{R: 99, G: 99, B: 99, A: 255},
	ScrollBarColor:      color.RGBA{R: 35, G: 35, B: 35, A: 8},
	ShadowColor:         color.RGBA{R: 0, G: 0, B: 0, A: 64},
	TextSize:            13,
	TextFont:            theme.DarkTheme().Font(fyne.TextStyle{}),
	TextBoldFont:        theme.DarkTheme().Font(fyne.TextStyle{Bold: true}),
	TextItalicFont:      theme.DarkTheme().Font(fyne.TextStyle{Italic: true}),
	TextBoldItalicFont:  theme.DarkTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
	TextMonospaceFont:   theme.DarkTheme().Font(fyne.TextStyle{Monospace: true}),
	Padding:             4,
	IconInlineSize:      24,
	ScrollBarSize:       10,
	ScrollBarSmallSize:  4,
}

var ExampleTree = PrettyTheme{
	BackgroundColor:     &color.RGBA{R: 0x70, G: 0x78, B: 0x4c, A: 0xf0},
	ButtonColor:         &color.RGBA{R: 0x70, G: 0x78, B: 0x4c, A: 0xd2},
	DisabledButtonColor: &color.RGBA{R: 0x70, G: 0x78, B: 0x4c, A: 0xff},
	HyperlinkColor:      &color.RGBA{R: 0x1, G: 0x24, B: 0x1, A: 0xff},
	TextColor:           &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	DisabledTextColor:   &color.RGBA{R: 0x2a, G: 0x32, B: 0x6, A: 0xff},
	IconColor:           &color.RGBA{R: 0x3c, G: 0x34, B: 0xf, A: 0xff},
	DisabledIconColor:   &color.RGBA{R: 0x5c, G: 0x64, B: 0x38, A: 0xff},
	PlaceHolderColor:    &color.RGBA{R: 0x57, G: 0x2c, B: 0x1, A: 0xff},
	PrimaryColor:        &color.RGBA{R: 0x70, G: 0x78, B: 0x4c, A: 0xb4},
	HoverColor:          &color.RGBA{R: 0x70, G: 0x78, B: 0x4c, A: 0xc8},
	FocusColor:          &color.RGBA{R: 0xa0, G: 0x96, B: 0x32, A: 0xff},
	ScrollBarColor:      &color.RGBA{R: 0x23, G: 0x23, B: 0x23, A: 0x8},
	ShadowColor:         &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x40},
	TextSize:            15,
	TextFont:            theme.DarkTheme().Font(fyne.TextStyle{}),
	TextBoldFont:        theme.DarkTheme().Font(fyne.TextStyle{Bold: true}),
	TextItalicFont:      theme.DarkTheme().Font(fyne.TextStyle{Italic: true}),
	TextBoldItalicFont:  theme.DarkTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
	TextMonospaceFont:   theme.DarkTheme().Font(fyne.TextStyle{Monospace: true}),
	Padding:             4,
	IconInlineSize:      28,
	ScrollBarSize:       10,
	ScrollBarSmallSize:  4,
}

var ExampleDracula = PrettyTheme{
	BackgroundColor:     &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	ButtonColor:         &color.RGBA{R: 0xf, G: 0xf, B: 0xf, A: 0xff},
	DisabledButtonColor: &color.RGBA{R: 0xf, G: 0xf, B: 0x11, A: 0xff},
	HyperlinkColor:      &color.RGBA{R: 0x0, G: 0x41, B: 0x78, A: 0xff},
	TextColor:           &color.RGBA{R: 0xaf, G: 0xaf, B: 0xaf, A: 0xff},
	DisabledTextColor:   &color.RGBA{R: 0x4e, G: 0x4e, B: 0x4e, A: 0xff},
	IconColor:           &color.RGBA{R: 0x53, G: 0x0, B: 0x0, A: 0xff},
	DisabledIconColor:   &color.RGBA{R: 0x21, G: 0x21, B: 0x21, A: 0xff},
	PlaceHolderColor:    &color.RGBA{R: 0x64, G: 0x14, B: 0x27, A: 0xff},
	PrimaryColor:        &color.RGBA{R: 0x50, G: 0x0, B: 0x13, A: 0xff},
	HoverColor:          &color.RGBA{R: 0x1a, G: 0x12, B: 0xe, A: 0xff},
	FocusColor:          &color.RGBA{R: 0x28, G: 0x0, B: 0x9, A: 0x80},
	ScrollBarColor:      &color.RGBA{R: 0x23, G: 0x23, B: 0x23, A: 0x8},
	ShadowColor:         &color.RGBA{R: 0x12, G: 0x12, B: 0x12, A: 0xff},
	TextSize:            16,
	TextFont:            theme.DarkTheme().Font(fyne.TextStyle{}),
	TextBoldFont:        theme.DarkTheme().Font(fyne.TextStyle{Bold: true}),
	TextItalicFont:      theme.DarkTheme().Font(fyne.TextStyle{Italic: true}),
	TextBoldItalicFont:  theme.DarkTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
	TextMonospaceFont:   theme.DarkTheme().Font(fyne.TextStyle{Monospace: true}),
	Padding:             3,
	IconInlineSize:      23,
	ScrollBarSize:       10,
	ScrollBarSmallSize:  4,
}

var ExampleCubicleLife = PrettyTheme{
	BackgroundColor:     &color.RGBA{R: 0xd3, G: 0xd3, B: 0xd3, A: 0xff},
	ButtonColor:         &color.RGBA{R: 0xbf, G: 0xbf, B: 0xbf, A: 0xff},
	DisabledButtonColor: &color.RGBA{R: 0x87, G: 0x87, B: 0x87, A: 0xff},
	HyperlinkColor:      &color.RGBA{R: 0x0, G: 0x33, B: 0xff, A: 0xff},
	TextColor:           &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	DisabledTextColor:   &color.RGBA{R: 0x1a, G: 0x1a, B: 0x1a, A: 0xff},
	IconColor:           &color.RGBA{R: 0x25, G: 0x25, B: 0x25, A: 0xff},
	DisabledIconColor:   &color.RGBA{R: 0xa9, G: 0xa9, B: 0xa9, A: 0xff},
	PlaceHolderColor:    &color.RGBA{R: 0x94, G: 0x94, B: 0x94, A: 0xff},
	PrimaryColor:        &color.RGBA{R: 0x33, G: 0x66, B: 0xcc, A: 0xf0},
	HoverColor:          &color.RGBA{R: 0x8f, G: 0x8f, B: 0x8f, A: 0xff},
	FocusColor:          &color.RGBA{R: 0xa7, G: 0xbd, B: 0xdf, A: 0xff},
	ScrollBarColor:      &color.RGBA{R: 0x4d, G: 0x4d, B: 0x4d, A: 0xff},
	ShadowColor:         &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x40},
	TextSize:            15,
	TextFont:            theme.DarkTheme().Font(fyne.TextStyle{}),
	TextBoldFont:        theme.DarkTheme().Font(fyne.TextStyle{Bold: true}),
	TextItalicFont:      theme.DarkTheme().Font(fyne.TextStyle{Italic: true}),
	TextBoldItalicFont:  theme.DarkTheme().Font(fyne.TextStyle{Bold: true, Italic: true}),
	TextMonospaceFont:   theme.DarkTheme().Font(fyne.TextStyle{Monospace: true}),
	Padding:             6,
	IconInlineSize:      32,
	ScrollBarSize:       10,
	ScrollBarSmallSize:  4,
}
