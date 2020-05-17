package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
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
	apply := &widget.Button{}

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

	bgc := screens.ColorSlider("BackgroundColor", pfc.BackgroundColor, BackgroundColorChan)
	bc := screens.ColorSlider("ButtonColor", pfc.ButtonColor, ButtonColorChan)
	dbc := screens.ColorSlider("DisabledButtonColor", pfc.DisabledButtonColor, DisabledButtonColorChan)
	hlc := screens.ColorSlider("HyperlinkColor", pfc.HyperlinkColor, HyperlinkColorChan)
	tc := screens.ColorSlider("TextColor", pfc.TextColor, TextColorChan)
	dtc := screens.ColorSlider("DisabledTextColor", pfc.DisabledTextColor, DisabledTextColorChan)
	ic := screens.ColorSlider("IconColor", pfc.IconColor, IconColorChan)
	dic := screens.ColorSlider("DisabledIconColor", pfc.DisabledIconColor, DisabledIconColorChan)
	phc := screens.ColorSlider("PlaceHolderColor", pfc.PlaceHolderColor, PlaceHolderColorChan)
	pc := screens.ColorSlider("PrimaryColor", pfc.PrimaryColor, PrimaryColorChan)
	hc := screens.ColorSlider("HoverColor", pfc.HoverColor, HoverColorChan)
	fc := screens.ColorSlider("FocusColor", pfc.FocusColor, FocusColorChan)
	sbc := screens.ColorSlider("ScrollBarColor", pfc.ScrollBarColor, ScrollBarColorChan)
	sc := screens.ColorSlider("ShadowColor", pfc.ShadowColor, ShadowColorChan)

	toRbga := func(c color.Color) *color.RGBA {
		r, g, b, a := c.RGBA()
		return &color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	}
	set := func(t prettyfyne.PrettyTheme) {
		BackgroundColorChan <- toRbga(t.BackgroundColor)
		ButtonColorChan <- toRbga(t.ButtonColor)
		DisabledButtonColorChan <- toRbga(t.DisabledButtonColor)
		HyperlinkColorChan <- toRbga(t.HyperlinkColor)
		TextColorChan <- toRbga(t.TextColor)
		DisabledTextColorChan <- toRbga(t.DisabledTextColor)
		IconColorChan <- toRbga(t.IconColor)
		DisabledIconColorChan <- toRbga(t.DisabledIconColor)
		PlaceHolderColorChan <- toRbga(t.PlaceHolderColor)
		PrimaryColorChan <- toRbga(t.PrimaryColor)
		HoverColorChan <- toRbga(t.HoverColor)
		FocusColorChan <- toRbga(t.FocusColor)
		ScrollBarColorChan <- toRbga(t.ScrollBarColor)
		ShadowColorChan <- toRbga(t.ShadowColor)
		bgc.Objects = screens.ColorSlider("BackgroundColor", toRbga(t.BackgroundColor), BackgroundColorChan).Objects
		bc.Objects = screens.ColorSlider("ButtonColor", toRbga(t.ButtonColor), ButtonColorChan).Objects
		dbc.Objects = screens.ColorSlider("DisabledButtonColor", toRbga(t.DisabledButtonColor), DisabledButtonColorChan).Objects
		hlc.Objects = screens.ColorSlider("HyperlinkColor", toRbga(t.HyperlinkColor), HyperlinkColorChan).Objects
		tc.Objects = screens.ColorSlider("TextColor", toRbga(t.TextColor), TextColorChan).Objects
		dtc.Objects = screens.ColorSlider("DisabledTextColor", toRbga(t.DisabledTextColor), DisabledTextColorChan).Objects
		ic.Objects = screens.ColorSlider("IconColor", toRbga(t.IconColor), IconColorChan).Objects
		dic.Objects = screens.ColorSlider("DisabledIconColor", toRbga(t.DisabledIconColor), DisabledIconColorChan).Objects
		phc.Objects = screens.ColorSlider("PlaceHolderColor", toRbga(t.PlaceHolderColor), PlaceHolderColorChan).Objects
		pc.Objects = screens.ColorSlider("PrimaryColor", toRbga(t.PrimaryColor), PrimaryColorChan).Objects
		hc.Objects = screens.ColorSlider("HoverColor", toRbga(t.HoverColor), HoverColorChan).Objects
		fc.Objects = screens.ColorSlider("FocusColor", toRbga(t.FocusColor), FocusColorChan).Objects
		sbc.Objects = screens.ColorSlider("ScrollBarColor", toRbga(t.ScrollBarColor), ScrollBarColorChan).Objects
		sc.Objects = screens.ColorSlider("ShadowColor", toRbga(t.ShadowColor), ShadowColorChan).Objects
		//w.Content().Refresh()
		apply.Tapped(&fyne.PointEvent{})
	}
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Import", fyne.NewMenuItem("Paste Yaml", func() {
			go func() {
				y := loadPastedYaml()
				if y != nil {
					set(*y)
				}
			}()
		})),
	))
	quickSelect := widget.NewSelect(
		[]string{
			"Default Theme",
			"Default Light Theme",
			"Every Day is Halloween",
			"Material Light",
			"Tree",
			"Dracula",
			"Cubicle Life",
		},
		func(s string) {
			switch s {
			case "Default Theme":
				set(*prettyfyne.DefaultTheme())
			case "Default Light Theme":
				set(*prettyfyne.DefaultLightTheme())
			case "Every Day is Halloween":
				set(prettyfyne.ExampleEveryDayIsHalloween)
			case "Material Light":
				set(prettyfyne.ExampleMaterialLight)
			case "Tree":
				set(prettyfyne.ExampleTree)
			case "Dracula":
				set(prettyfyne.ExampleDracula)
			case "Cubicle Life":
				set(prettyfyne.ExampleCubicleLife)
			}
		},
	)

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

	apply = widget.NewButtonWithIcon("Apply", theme.ConfirmIcon(), func(){
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
		widget.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Load Example:"),
			quickSelect,
			layout.NewSpacer(),
		),
		bgc,
		bc,
		dbc,
		hlc,
		tc,
		dtc,
		ic,
		dic,
		phc,
		pc,
		hc,
		fc,
		sbc,
		sc,
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

func loadPastedYaml() *prettyfyne.PrettyTheme {
	themeChan := make(chan *prettyfyne.PrettyTheme)
	cancelChan := make(chan bool)
	pop := &widget.PopUp{}
	winSize := fyne.CurrentApp().Driver().AllWindows()[0].Canvas().Size()
	in := widget.NewMultiLineEntry()
	cancel := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
		cancelChan <- true
		pop.Hide()
	})
	apply := widget.NewButtonWithIcon("Load", theme.ConfirmIcon(), func() {
		pt, _, err := prettyfyne.UnmarshalYaml([]byte(in.Text))
		if err != nil {
			go dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		pop.Hide()
		themeChan <- pt
	})

	pop = widget.NewPopUp(
		fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(winSize.Width,winSize.Height)),
			widget.NewVBox(
				fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(winSize.Width,winSize.Height-50)),
					widget.NewScrollContainer(in),
				),
				widget.NewHBox(layout.NewSpacer(), cancel, apply, layout.NewSpacer()),
			),
		),
		fyne.CurrentApp().Driver().AllWindows()[0].Canvas(),
	)

	pop.Show()
	for {
		select {
		case <-cancelChan:
			return nil
		case t := <-themeChan:
			return t
		}
	}

}
