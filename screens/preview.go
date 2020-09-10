package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func Preview() *fyne.Container {

	toolbar := widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func(){}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func(){}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func(){}),
		widget.NewToolbarAction(theme.ContentPasteIcon(),func(){}),
	)


	boldLabel := widget.NewLabelWithStyle("Bold Label", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
	})
	italicLabel := widget.NewLabelWithStyle("Italic Label", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      false,
		Italic:    true,
		Monospace: false,
	})
	boldItalicLabel := widget.NewLabelWithStyle("Bold Italic Label", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    true,
		Monospace: false,
	})
	monoLabel := widget.NewLabelWithStyle("Monospace Bold Label", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: true,
	})

	entryEnabled := widget.NewEntry()
	entryEnabled.SetPlaceHolder("Type Something Here.")
	entryDisabled := widget.NewEntry()
	entryDisabled.SetText("This entry is disabled")
	entryDisabled.Disable()
	entryHasText := widget.NewEntry()
	entryHasText.SetText("This entry already has text")

	buttonDisabled := widget.NewButtonWithIcon("Disabled", theme.VisibilityIcon(), func() {})
	buttonDisabled.Disable()

	slider := widget.NewSlider(0, 10)
	slider.Value = 3.0

	pass := widget.NewPasswordEntry()
	pass.SetText("P@55w0rd")

	progLabel := widget.NewLabel("waiting . . .")
	progress := widget.NewProgressBar()
	go func() {
		t := time.NewTicker(50 * time.Millisecond)
		i := float64(0)
		for {
			select {
			case <-t.C:
				i = i + 1
				progress.SetValue(i / 1000)
				if i >= 1000 {
					i = 1
					progLabel.SetText("Starting over ... please wait.")
					progLabel.TextStyle = fyne.TextStyle{
						Bold:      true,
						Italic:    true,
					}
				}
				switch {
				case i == 999:
					progLabel.SetText("Something happened. Starting over.")
					time.Sleep(3*time.Second)
				case i > 900:
					progLabel.SetText("Get Ready! Here it comes")
					progLabel.TextStyle = fyne.TextStyle{
						Bold:      true,
						Italic:    false,
					}
				case i > 500:
					progLabel.SetText("getting closer!")
				case i > 100:
					progLabel.SetText("la dee da, la dee da dee da ....")
					progLabel.TextStyle = fyne.TextStyle{
						Bold:      false,
						Italic:    false,
					}
				}
			}
		}
	}()

	timeRem := widget.NewLabelWithStyle("Time Remaining:                    1 minute", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	go func(){
		tr := uint64(1)
		p := message.NewPrinter(language.AmericanEnglish)
		for {
			time.Sleep(time.Duration(rand.Intn(2)+1)*time.Second)
			timeRem.SetText(p.Sprintf("Time Remaining: %19d minutes", tr))
			tr = tr + uint64(rand.Intn(5000))
		}
	}()

	multiLineRo := widget.NewMultiLineEntry()
	multiLineRo.Disable()
	go func(){
		var scrolling = make([]string, 10)
		rPush := func(s string) {
			scrolling = append(scrolling[1:], s)
		}
		for {
			for _, line := range scrollText {
				rPush(line)
				multiLineRo.SetText(strings.Join(scrolling, "\n"))
				time.Sleep(750*time.Millisecond)
			}
		}
	}()

	form := widget.NewForm()
	selection := widget.NewSelect([]string{"Yes", "Are you kidding?", "Sure, I'm on the clock", "No time for love Dr. Jones."}, func(s string){})
	selection.SetSelected("Are you kidding?")
	formEntrySelects := widget.NewFormItem("Should we wait?", widget.NewHBox(
		selection,
		widget.NewSelect([]string{"one"}, func(s string){}),
		layout.NewSpacer(),
	))
	form.AppendItem(formEntrySelects)

	radio := widget.NewRadio([]string{"Yes", "Not yet", "No thanks, it's 3am!", "I would rather have something else."}, func(s string){})
	form.Append("Coffee?", widget.NewHBox(radio, layout.NewSpacer()))

	alreadyChecked := widget.NewCheck("Already Checked this one.", func(b bool){})
	alreadyChecked.SetChecked(true)
	scrollGroup := widget.NewGroupWithScroller("Check these out",
		widget.NewCheck("Having fun?", func(b bool){}),
		widget.NewCheck("How about now?", func(b bool){}),
		alreadyChecked,
		widget.NewCheck("Here is another check", func(b bool){}),
		widget.NewCheck("Do these go on forever?", func(b bool){}),
		widget.NewCheck("Last one.", func(b bool){}),
	)

	uri, _ := url.Parse("https://github.com/blockpane/prettyfyne")

	return fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		widget.NewScrollContainer(
			widget.NewVBox(
				toolbar,
				widget.NewHBox(widget.NewHyperlink("Link to prettyfyne", uri), layout.NewSpacer(), boldLabel, layout.NewSpacer(), boldItalicLabel, layout.NewSpacer(), widget.NewLabel("This is a regular label."), layout.NewSpacer(), italicLabel),
				layout.NewSpacer(),
				fyne.NewContainerWithLayout(layout.NewGridLayout(3), slider, widget.NewIcon(theme.SettingsIcon()),
					fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(320, 150)), scrollGroup),
				),
				layout.NewSpacer(),
				widget.NewHBox(layout.NewSpacer(), entryEnabled, pass, entryDisabled, entryHasText,layout.NewSpacer()),
				layout.NewSpacer(),
				fyne.NewContainerWithLayout(layout.NewGridLayout(3), layout.NewSpacer(), progress, progLabel),
				layout.NewSpacer(),
				widget.NewHBox(layout.NewSpacer(), widget.NewButtonWithIcon("Enabled", theme.InfoIcon(), func() {}), monoLabel, buttonDisabled, layout.NewSpacer()),
				layout.NewSpacer(),
				widget.NewHBox(layout.NewSpacer(), timeRem, layout.NewSpacer()),
				widget.NewProgressBarInfinite(),
				form,
				layout.NewSpacer(),
				widget.NewHBox(layout.NewSpacer(),
					fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(600, 300)),
						multiLineRo,
					),
					layout.NewSpacer(),
				),
			),
		),
	)
}

var scrollText = [...]string{
	"Getting Started",
	"",
	"Using the Fyne toolkit to build cross platform applications is very simple ",
	"but does require some tools to be installed before you can begin.",
	"If your computer is set up for development with Go then the following steps may not be required, ",
	"but we advise reading the tips for your operating system just in case.",
	"If later steps in this tutorial fail then you should re-visit the prerequisites below.",
	"Prerequisites",
	"",
	"Fyne requires 3 basic elements to be present, the Go tools (at least version 1.12), ",
	"a C compiler (to connect with system graphics drivers) and an system graphics driver.",
	"The instructions vary depending on your operating system, ",
	"choose the appropriate tab below for installation instructions.",
	"",
	"Note that these steps are just required for development - ",
	"your Fyne applications will not require any setup or dependency installation for end users!",
	"",
	"Downloading",
	"After installing any prerequisites the following command will do everything to get Fyne installed:",
	"",
	"$ go get fyne.io/fyne",
	"",
	"Once that command completes you will have the full Fyne development package installed in your GOPATH.",
	"",
	"Run demo",
	"",
	"If you want to see the Fyne toolkit in action before you start to code your own application ",
	"then you can see our demo app running on your computer by executing:",
	"",
	"$ go get fyne.io/fyne/cmd/fyne_demo",
	"$ fyne_demo",
	"",
	"And that’s all there is to it! Now you can write your own Fyne application in your IDE of choice.",
	"",
}

