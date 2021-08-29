package prettyfyne

import (
	"image/color"

	"gopkg.in/yaml.v2"
)

// PrettyThemeConfig is used for serialization and loading from yaml
type PrettyThemeConfig struct {
	BackgroundColor        *color.RGBA `yaml:"background_color,omitempty"`
	ButtonColor            *color.RGBA `yaml:"button_color,omitempty"`
	DisabledButtonColor    *color.RGBA `yaml:"disabled_button_color,omitempty"`
	HyperlinkColor         *color.RGBA `yaml:"hyperlink_color,omitempty"`
	TextColor              *color.RGBA `yaml:"text_color,omitempty"`
	DisabledTextColor      *color.RGBA `yaml:"disabled_text_color,omitempty"`
	IconColor              *color.RGBA `yaml:"icon_color,omitempty"`
	DisabledIconColor      *color.RGBA `yaml:"disabled_icon_color,omitempty"`
	PlaceHolderColor       *color.RGBA `yaml:"place_holder_color,omitempty"`
	PrimaryColor           *color.RGBA `yaml:"primary_color,omitempty"`
	HoverColor             *color.RGBA `yaml:"hover_color,omitempty"`
	FocusColor             *color.RGBA `yaml:"focus_color,omitempty"`
	ScrollBarColor         *color.RGBA `yaml:"scroll_bar_color,omitempty"`
	ShadowColor            *color.RGBA `yaml:"shadow_color,omitempty"`
	TextSize               float32     `yaml:"text_size,omitempty"`
	TextFontPath           string      `yaml:"text_font_path,omitempty"`
	TextFont               string      `yaml:"text_font,omitempty"`
	TextBoldFontPath       string      `yaml:"text_bold_font_path,omitempty"`
	TextBoldFont           string      `yaml:"text_bold_font,omitempty"`
	TextItalicFontPath     string      `yaml:"text_italic_font_path,omitempty"`
	TextItalicFont         string      `yaml:"text_italic_font,omitempty"`
	TextBoldItalicFontPath string      `yaml:"text_bold_italic_font_path,omitempty"`
	TextBoldItalicFont     string      `yaml:"text_bold_italic_font,omitempty"`
	TextMonospaceFontPath  string      `yaml:"text_monospace_font_path,omitempty"`
	TextMonospaceFont      string      `yaml:"text_monospace_font,omitempty"`
	Padding                float32     `yaml:"padding,omitempty"`
	IconInlineSize         float32     `yaml:"icon_inline_size,omitempty"`
	ScrollBarSize          float32     `yaml:"scroll_bar_size,omitempty"`
	ScrollBarSmallSize     float32     `yaml:"scroll_bar_small_size,omitempty"`
}

// UnmarshalYaml will override the default theme settings with what is stored in a yaml file. All fields are optional
// and will fall back to the default if not set. It will always return a populated theme, even if it cannot load fonts.
func UnmarshalYaml(y []byte) (pt *PrettyTheme, fontsLoaded bool, err error) {
	pt = DefaultTheme()
	c := PrettyThemeConfig{}
	err = yaml.Unmarshal(y, &c)
	if err != nil {
		return
	}
	fontsLoaded = true
	//switch {
	if c.BackgroundColor != nil {
		pt.BackgroundColor = c.BackgroundColor
	}
	if c.ButtonColor != nil {
		pt.ButtonColor = c.ButtonColor
	}
	if c.DisabledButtonColor != nil {
		pt.DisabledButtonColor = c.DisabledButtonColor
	}
	if c.HyperlinkColor != nil {
		pt.HyperlinkColor = c.HyperlinkColor
	}
	if c.TextColor != nil {
		pt.TextColor = c.TextColor
	}
	if c.DisabledTextColor != nil {
		pt.DisabledTextColor = c.DisabledTextColor
	}
	if c.IconColor != nil {
		pt.IconColor = c.IconColor
	}
	if c.DisabledIconColor != nil {
		pt.DisabledIconColor = c.DisabledIconColor
	}
	if c.PlaceHolderColor != nil {
		pt.PlaceHolderColor = c.PlaceHolderColor
	}
	if c.PrimaryColor != nil {
		pt.PrimaryColor = c.PrimaryColor
	}
	if c.HoverColor != nil {
		pt.HoverColor = c.HoverColor
	}
	if c.FocusColor != nil {
		pt.FocusColor = c.FocusColor
	}
	if c.ScrollBarColor != nil {
		pt.ScrollBarColor = c.ScrollBarColor
	}
	if c.ShadowColor != nil {
		pt.ShadowColor = c.ShadowColor
	}
	if c.TextSize != 0 {
		pt.TextSize = c.TextSize
	}
	if c.Padding != 0 {
		pt.Padding = c.Padding
	}
	if c.IconInlineSize != 0 {
		pt.IconInlineSize = c.IconInlineSize
	}
	if c.ScrollBarSize != 0 {
		pt.ScrollBarSize = c.ScrollBarSize
	}
	if c.ScrollBarSmallSize != 0 {
		pt.ScrollBarSmallSize = c.ScrollBarSmallSize
	}
	// handle loading of fonts
	if c.TextFontPath != "" || c.TextFont != "" {
		pt.TextFont, err = LoadFont(c.TextFontPath, c.TextFont, "NotoSans-Regular.ttf")
		if err != nil {
			fontsLoaded = false
		}
	}
	if c.TextBoldFontPath != "" || c.TextBoldFont != "" {
		pt.TextBoldFont, err = LoadFont(c.TextBoldFontPath, c.TextBoldFont, "NotoSans-Bold.ttf")
		if err != nil {
			fontsLoaded = false
		}
	}
	if c.TextBoldItalicFontPath != "" || c.TextBoldItalicFont != "" {
		pt.TextBoldItalicFont, err = LoadFont(c.TextBoldItalicFontPath, c.TextBoldItalicFont, "NotoSans-BoldItalic.ttf")
		if err != nil {
			fontsLoaded = false
		}
	}
	if c.TextItalicFontPath != "" || c.TextItalicFont != "" {
		pt.TextItalicFont, err = LoadFont(c.TextItalicFontPath, c.TextItalicFont, "NotoSans-Italic.ttf")
		if err != nil {
			fontsLoaded = false
		}
	}
	if c.TextMonospaceFontPath != "" || c.TextMonospaceFont != "" {
		pt.TextMonospaceFont, err = LoadFont(c.TextMonospaceFontPath, c.TextMonospaceFont, "NotoMono-Regular.ttf")
		if err != nil {
			fontsLoaded = false
		}
	}
	return pt, fontsLoaded, nil
}
