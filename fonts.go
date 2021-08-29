package prettyfyne

import (
	"bufio"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
)

type Font string

const (
	FontDefault    Font = "NotoSans-Regular.ttf"
	FontBold       Font = "NotoSans-Bold.ttf"
	FontItalic     Font = "NotoSans-Italic.ttf"
	FontBoldItalic Font = "NotoSans-BoldItalic.ttf"
	FontMono       Font = "NotoMono-Regular.ttf"
)

// LoadFont gives a best-effort attempt at loading a font, and returns the default font if it fails
func LoadFont(path string, name string, fallbackFont Font) (font fyne.Resource, err error) {
	if fallbackFont == "" {
		fallbackFont = FontDefault
	}

	getFallback := func(name Font) fyne.Resource {
		switch name {
		case FontDefault:
			return theme.DefaultTextFont()
		case FontBold:
			return theme.DefaultTextBoldFont()
		case FontBoldItalic:
			return theme.DefaultTextBoldItalicFont()
		case FontItalic:
			return theme.DefaultTextItalicFont()
		case FontMono:
			return theme.DefaultTextMonospaceFont()
		}
		return nil
	}

	if path == "" && name == "" {
		return getFallback(fallbackFont), errors.New("no font name or path provided")
	}
	// quick check if it's one of the default fyne theme fonts
	font = getFallback(Font(name))
	if font != nil {
		return
	}

	// see if it's a system font
	if path == "" {
		sysFonts, err := GetFonts()
		if err != nil {
			return getFallback(fallbackFont), err
		}

		// a few locations to look on mac:
		if sysFonts[name] == "/Library/Fonts" {
			if _, err := os.Lstat( "/System/Library/Fonts/Supplemental/"+name); err == nil {
				path = "/System/Library/Fonts/Supplemental"
			}
		} else if _, err := os.Lstat( "/System/Library/Fonts/"+name); err == nil {
			path = "/System/Library/Fonts"
		} else {
			path = sysFonts[name]
		}
		if path == "" {
			return getFallback(fallbackFont), errors.New("could not find requested font, not in known fonts path")
		}
	}
	fullPath, _ := regexp.MatchString(`(?i)tt[fc]`, path) // allow putting full path to font in path field
	if !strings.HasSuffix(path, string(os.PathSeparator)) && !fullPath {
		path = path + string(os.PathSeparator)
	}
	font, err = fyne.LoadResourceFromPath(path + name)
	if err != nil {
		return getFallback(fallbackFont), err
	}
	if font != nil {
		return
	}
	return getFallback(fallbackFont), errors.New("could not load font")
}

// GetFonts provides a map containing a list of fonts and their location based on system defaults,
// TODO: for now only works on MacOS
func GetFonts() (fonts map[string]string, err error) {
	fonts = make(map[string]string)
	switch runtime.GOOS {
	case "darwin":
		defaultFonts := []string{"/opt/X11/share/fonts/TTF/fonts.list", "/Library/Fonts/fonts.list"}
		if u, e := user.Current(); e == nil && u.HomeDir != "" {
			defaultFonts = append(defaultFonts, u.HomeDir+"/Library/Fonts/fonts.list")
		}
		for _, path := range defaultFonts {
			f, err := os.Open(path)
			if err == nil {
				func() {
					defer f.Close()
					reader := bufio.NewReader(f)
					for {
						line, err := reader.ReadString('\n')
						if err != nil {
							return
						}
						if match, _ := regexp.MatchString(`(?i)tt[fc]`, line); match {
							entry := strings.Split(line, "//")
							if len(entry) == 2 {
								fonts[strings.TrimSpace(entry[1])] = entry[0]
							}
						}
					}
				}()
			}
		}
		if len(fonts) == 0 {
			return fonts, errors.New("none found")
		}
		return
	}
	return fonts, errors.New("don't know the default fonts path for " + runtime.GOOS)
}
