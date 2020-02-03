package screens

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/frameloss/prettyfyne"
	"image"
	"image/color"
	"sort"
	"strconv"
	"sync"
	"time"
)

func ColorSlider(label string, rgba *color.RGBA, newColor chan *color.RGBA) *fyne.Container {
	// closure to update rgba pointer
	r := func(c color.Color) {
		r, g, b, a := c.RGBA()
		rgba = &color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	}
	// closure to dereference so we have a copy of original to revert to
	oldRgba := func(c *color.RGBA) color.RGBA {
		return *c
	}(rgba)
	// closure to copy the dereferenced original back
	resetRgba := func() *color.RGBA {
		c := oldRgba
		return &c
	}

	updated := func() {
		newColor <- rgba
	}

	palletMap := prettyfyne.PalletMap()
	palletMap["Default"] = []color.Color{rgba}
	colorPallet := palletMap["Default"]
	pallets := func() []string {
		s := []string{"Values"}
		for p := range palletMap {
			s = append(s, p)
		}
		sort.Strings(s)
		return s
	}()


	spacer := layout.NewSpacer()
	sampleColor := canvas.NewImageFromImage(imageRgba(rgba, 36, 36))

	rEntry := widget.NewEntry()
	rMux := sync.Mutex{}
	rEntry.OnChanged = func(s string) {
		rgba.R = checkRgbaInput(s)
		go func() {
			rMux.Lock()
			time.Sleep(300 * time.Millisecond)
			if rgba.R == 0 && rEntry.Text != "0" {
				rEntry.SetText("0")
			}
			rMux.Unlock()
		}()
		sampleColor.Image = imageRgba(rgba , 36, 36)
		sampleColor.Refresh()
		updated()
	}

	gEntry := widget.NewEntry()
	gMux := sync.Mutex{}
	gEntry.OnChanged = func(s string) {
		rgba.G = checkRgbaInput(s)
		go func() {
			gMux.Lock()
			time.Sleep(300 * time.Millisecond)
			if rgba.G == 0 && gEntry.Text != "0" {
				gEntry.SetText("0")
			}
			gMux.Unlock()
		}()
		sampleColor.Image = imageRgba(rgba , 36, 36)
		sampleColor.Refresh()
		updated()
	}

	bEntry := widget.NewEntry()
	bMux := sync.Mutex{}
	bEntry.OnChanged = func(s string) {
		rgba.B = checkRgbaInput(s)
		go func() {
			bMux.Lock()
			time.Sleep(300 * time.Millisecond)
			if rgba.B == 0 && bEntry.Text != "0" {
				bEntry.SetText("0")
			}
			bMux.Unlock()
		}()
		sampleColor.Image = imageRgba(rgba , 36, 36)
		sampleColor.Refresh()
		updated()
	}

	aEntry := widget.NewEntry()
	aMux := sync.Mutex{}
	aEntry.OnChanged = func(s string) {
		rgba.A = checkRgbaInput(s)
		go func() {
			aMux.Lock()
			time.Sleep(300 * time.Millisecond)
			if rgba.A == 0 && aEntry.Text != "0" {
				aEntry.SetText("0")
			}
			aMux.Unlock()
		}()
		sampleColor.Image = imageRgba(rgba , 36, 36)
		sampleColor.Refresh()
		updated()
	}

	rgbaEntries := widget.NewHBox(rEntry, gEntry, bEntry, aEntry)

	slider := widget.NewSlider(0, float64(len(colorPallet)-1))
	slider.Step = 1.0
	slider.OnChanged = func(f float64) {
		r(colorPallet[int(f)])
		sampleColor.Image = imageRgba(rgba , 36, 36)
		sampleColor.Refresh()
		updated()
	}
	colorSelect := widget.NewSelect(pallets, func(s string){
		switch s {
		case "Default":
			slider.Hide()
			rgba = resetRgba()
			rgbaEntries.Hide()
			spacer.Show()
		case "Values":
			slider.Hide()
			spacer.Hide()
			rEntry.SetText(fmt.Sprint(rgba.R))
			gEntry.SetText(fmt.Sprint(rgba.G))
			bEntry.SetText(fmt.Sprint(rgba.B))
			aEntry.SetText(fmt.Sprint(rgba.A))
			rgbaEntries.Show()
		default:
			slider.Show()
			colorPallet = palletMap[s]
			slider.Max = float64(len(palletMap[s])-1)
			slider.Value = float64(len(palletMap[s])/2)
			r(colorPallet[int(slider.Value)])
			spacer.Hide()
			rgbaEntries.Hide()
		}
		sampleColor.Image = imageRgba(rgba , 36, 36)
		sampleColor.Refresh()
		slider.Refresh()
		updated()
	})
	colorSelect.SetSelected("Default")

	sampleColor.Refresh()
	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(4),
		widget.NewLabel(label),
		widget.NewHBox(
			widget.NewLabel("Colors"),
			colorSelect,
		),
		spacer,
		rgbaEntries,
		slider,
		sampleColor,
	)
}

func imageRgba(color *color.RGBA, h int, w int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color)
		}
	}
	return img
}

func checkRgbaInput(s string) uint8 {
	i, e := strconv.Atoi(s)
	if e != nil || i < 0 || i > 255 {
		return 0
	}
	return uint8(i)
}

