package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/frameloss/prettyfyne"
	"github.com/frameloss/prettyfyne/screens"
	"gopkg.in/yaml.v2"
	"image/color"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Theme Editor")

	pfc := prettyfyne.ExampleEveryDayIsHalloween.ToConfig()

	BackgroundColorChan := make(chan *color.RGBA)
	ButtonColorChan := make(chan *color.RGBA)
	DisabledButtonColorChan := make(chan *color.RGBA)
	HyperlinkColorChan := make(chan *color.RGBA)
	TextColorChan := make(chan *color.RGBA)
	DisabledTextColorChan := make(chan *color.RGBA)
	IconColorChan := make(chan *color.RGBA)
	DisabledIconColorChan := make(chan *color.RGBA)
	PlaceHolderColorChan := make(chan *color.RGBA)
	PrimaryColorChan := make(chan *color.RGBA)
	HoverColorChan := make(chan *color.RGBA)
	FocusColorChan := make(chan *color.RGBA)
	ScrollBarColorChan := make(chan *color.RGBA)
	ShadowColorChan := make(chan *color.RGBA)

	go func() {
		for {
			select {
			case c := <- BackgroundColorChan:
				pfc.BackgroundColor = c
			case c := <- ButtonColorChan:
				pfc.ButtonColor = c
			case c := <- DisabledButtonColorChan:
				pfc.DisabledButtonColor = c
			case c := <- HyperlinkColorChan:
				pfc.HyperlinkColor = c
			case c := <- TextColorChan:
				pfc.TextColor = c
			case c := <- DisabledTextColorChan:
				pfc.DisabledTextColor = c
			case c := <- IconColorChan:
				pfc.IconColor = c
			case c := <- DisabledIconColorChan:
				pfc.DisabledIconColor = c
			case c := <- PlaceHolderColorChan:
				pfc.PlaceHolderColor = c
			case c := <- PrimaryColorChan:
				pfc.PrimaryColor = c
			case c := <- HoverColorChan:
				pfc.HoverColor = c
			case c := <- FocusColorChan:
				pfc.FocusColor = c
			case c := <- ScrollBarColorChan:
				pfc.ScrollBarColor = c
			case c := <- ShadowColorChan:
				pfc.ShadowColor = c
			}
		}
	}()

	paddingValue := widget.NewLabel(fmt.Sprintf("%d", pfc.Padding))
	paddingSlider := widget.NewSlider(1.0, 32.0)
	paddingSlider.Step = 1.0
	paddingSlider.OnChanged = func(f float64) {
		paddingValue.SetText(fmt.Sprintf("%d", int(f)))
		pfc.Padding = int(f)
	}
	paddingSlider.Value = float64(pfc.Padding)
	padding := fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		widget.NewLabel("Padding"),
		layout.NewSpacer(),
		paddingSlider,
		paddingValue,
	)

	iconValue := widget.NewLabel(fmt.Sprintf("%d", pfc.IconInlineSize))
	iconSlider := widget.NewSlider(1, 128)
	iconSlider.Step = 1
	iconSlider.Value = float64(pfc.IconInlineSize)
	iconSlider.OnChanged = func(f float64) {
		iconValue.SetText(fmt.Sprintf("%d", int(f)))
		pfc.IconInlineSize = int(f)
	}
	icon := fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		widget.NewLabel("Icon Size"),
		layout.NewSpacer(),
		iconSlider,
		iconValue,
	)

	textSizeValue := widget.NewLabel(fmt.Sprintf("%d", pfc.TextSize))
	textSizeSlider := widget.NewSlider(1, 64)
	textSizeSlider.Step = 1
	textSizeSlider.Value = float64(pfc.TextSize)
	textSizeSlider.OnChanged = func(f float64) {
		textSizeValue.SetText(fmt.Sprintf("%d", int(f)))
		pfc.TextSize = int(f)
	}
	textSize := fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		widget.NewLabel("Text Size"),
		layout.NewSpacer(),
		textSizeSlider,
		textSizeValue,
	)

	//TODO: font selection
	/*
	textFontEntry := widget.NewEntry()
	textFontPathEntry := widget.NewEntry()
	textBoldFontEntry := widget.NewEntry()
	textBoldFontPathEntry := widget.NewEntry()
	textItalicFontEntry := widget.NewEntry()
	textItalicFontPathEntry := widget.NewEntry()
	textBoldItalicFontEntry := widget.NewEntry()
	textBoldItalicFontPathEntry := widget.NewEntry()
	textMonospaceFontEntry := widget.NewEntry()
	textMonospaceFontPathEntry := widget.NewEntry()
	 */

	tabSourceText := widget.NewMultiLineEntry()
	tabSourceText.SetText("Go source example will show here after applying the new theme")
	tabPftSourceText := widget.NewMultiLineEntry()
	tabPftSourceText.SetText("Go source example will show here after applying the new theme")
	tabYamlText := widget.NewMultiLineEntry()
	tabYamlText.SetText("YAML theme config will show here after applying the new theme")

	apply := widget.NewButtonWithIcon("Apply", theme.ConfirmIcon(), func(){
		y, _ := yaml.Marshal(pfc)
		pt, _, _ := prettyfyne.UnmarshalYaml(y)
		tabYamlText.SetText(string(y))
		tabSourceText.SetText(screens.ToGoSource(pfc))
		tabPftSourceText.SetText(screens.ToEasierGoSource(pfc))
		fyne.CurrentApp().Settings().SetTheme(pt.ToFyneTheme())
		w.Content().Refresh()
	})

	tabChooserContent := widget.NewScrollContainer(widget.NewVBox(
		widget.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Pretty Fyne Theme Editor"),
			layout.NewSpacer(),
		),
		screens.ColorSlider("BackgroundColor", pfc.BackgroundColor, BackgroundColorChan),
		screens.ColorSlider("ButtonColor", pfc.ButtonColor, ButtonColorChan),
		screens.ColorSlider("DisabledButtonColor", pfc.DisabledButtonColor, DisabledButtonColorChan),
		screens.ColorSlider("HyperlinkColor", pfc.HyperlinkColor, HyperlinkColorChan),
		screens.ColorSlider("TextColor", pfc.TextColor, TextColorChan),
		screens.ColorSlider("DisabledTextColor", pfc.DisabledTextColor, DisabledTextColorChan),
		screens.ColorSlider("IconColor", pfc.IconColor, IconColorChan),
		screens.ColorSlider("DisabledIconColor", pfc.DisabledIconColor, DisabledIconColorChan),
		screens.ColorSlider("PlaceHolderColor", pfc.PlaceHolderColor, PlaceHolderColorChan),
		screens.ColorSlider("PrimaryColor", pfc.PrimaryColor, PrimaryColorChan),
		screens.ColorSlider("HoverColor", pfc.HoverColor, HoverColorChan),
		screens.ColorSlider("FocusColor", pfc.FocusColor, FocusColorChan),
		screens.ColorSlider("ScrollBarColor", pfc.ScrollBarColor, ScrollBarColorChan),
		screens.ColorSlider("ShadowColor", pfc.ShadowColor, ShadowColorChan),
		padding,
		icon,
		textSize,
		widget.NewHBox(layout.NewSpacer(), apply, widget.NewButtonWithIcon("Exit", theme.CancelIcon(), func() {a.Quit()}), layout.NewSpacer()),
	))
	tabChooser := widget.NewTabItem("Settings", tabChooserContent)

	tabYamlContent := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		widget.NewScrollContainer(tabYamlText),
	)
	tabYaml := widget.NewTabItem("Generated Yaml", tabYamlContent)

	tabSourceContent := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		widget.NewScrollContainer(tabSourceText),
	)
	tabSource := widget.NewTabItem("Generated Go Source (Native Fyne Theme)", tabSourceContent)

	tabPftSourceContent := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		widget.NewScrollContainer(tabPftSourceText),
	)
	tabPftSource := widget.NewTabItem("Generated Go Source (Pretty Fyne Theme)", tabPftSourceContent)

	tabPreview := widget.NewTabItem("Preview Widgets", screens.Preview())

	tabs := widget.NewTabContainer(tabChooser, tabPreview, tabYaml, tabSource, tabPftSource)
	w.SetContent(
		fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(1000, 1000)),
			tabs,
		))
	go func(){
		for {
			time.Sleep(20 * time.Millisecond)
			if tabs.Visible() {
				break
			}
		}
		fyne.CurrentApp().Settings().SetTheme(prettyfyne.ExampleEveryDayIsHalloween.ToFyneTheme())
	}()
	w.CenterOnScreen()
	w.ShowAndRun()
}

