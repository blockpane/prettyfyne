package prettyfyne

import (
	"fmt"
	"testing"
)

func TestPrettyTheme_MarshalYaml(t *testing.T) {
	pt := DefaultTheme()
	y, err := pt.MarshalYaml()
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(y))
	}
}

func TestUnmarshalYaml(t *testing.T) {
	yt := `background_color:
  r: 66
  g: 66
  b: 66
  a: 255
button_color:
  r: 33
  g: 33
  b: 33
  a: 255
disabled_button_color:
  r: 49
  g: 49
  b: 49
  a: 255
hyperlink_color:
  r: 153
  g: 153
  b: 255
  a: 255
text_color:
  r: 255
  g: 255
  b: 255
  a: 255
disabled_text_color:
  r: 96
  g: 96
  b: 96
  a: 255
icon_color:
  r: 255
  g: 255
  b: 255
  a: 255
disabled_icon_color:
  r: 96
  g: 96
  b: 96
  a: 255
place_holder_color:
  r: 178
  g: 178
  b: 178
  a: 255
primary_color:
  r: 26
  g: 35
  b: 126
  a: 255
hover_color:
  r: 49
  g: 49
  b: 49
  a: 255
focus_color:
  r: 26
  g: 35
  b: 126
  a: 255
scroll_bar_color:
  r: 0
  g: 0
  b: 0
  a: 153
shadow_color:
  r: 0
  g: 0
  b: 0
  a: 102
text_size: 14
text_font: NotoSans-Regular.ttf
text_bold_font: NotoSans-Bold.ttf
text_italic_font: NotoSans-Italic.ttf
text_bold_italic_font: NotoSans-BoldItalic.ttf
text_monospace_font: NotoMono-Regular.ttf
padding: 4
icon_inline_size: 20
scroll_bar_size: 16
scroll_bar_small_size: 3
`
	_, ok, err := UnmarshalYaml([]byte(yt))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("didn't load fonts!")
	}
}

