package screens

import (
	"fmt"
	"github.com/frameloss/prettyfyne"
	"testing"
)

func TestToGoSource(t *testing.T) {
	fmt.Println(ToGoSource(prettyfyne.DefaultTheme().ToConfig()))
}
