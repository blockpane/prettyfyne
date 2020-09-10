package screens

import (
	"fmt"
	"github.com/blockpane/prettyfyne"
	"testing"
)

func TestToGoSource(t *testing.T) {
	fmt.Println(ToGoSource(prettyfyne.DefaultTheme().ToConfig()))
}
