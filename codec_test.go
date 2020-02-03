package prettyfyne

import (
	"runtime"
	"testing"
)

func TestGetFonts(t *testing.T) {
	if runtime.GOOS == "darwin" {
		if _, e := GetFonts(); e != nil {
			t.Error(e)
		}
	}
}
