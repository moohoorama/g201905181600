package font

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/bitmapfont"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const LineHeight = 16

func tokens(str string) []string {
	tokens := []string{}
	for {
		var strs []string
		switch len(tokens) % 2 {
		case 0:
			strs = strings.SplitN(str, "<red>", 2)
		case 1:
			strs = strings.SplitN(str, "</red>", 2)
		}
		if len(strs) >= 1 {
			tokens = append(tokens, strs[0])
		}
		if len(strs) == 2 {
			str = strs[1]
		} else {
			break
		}
	}
	return tokens
}

func Width(str string) int {
	w := fixed.I(0)
	for _, t := range tokens(str) {
		w += font.MeasureString(bitmapfont.Gothic12r, t)
	}
	return w.Round()
}

func Height(str string) int {
	return (strings.Count(str, "\n") + 1) * LineHeight
}

var red = color.RGBA{0xe4, 0x32, 0x60, 0xff}

func DrawText(target *ebiten.Image, str string, x, y int, clr color.Color) {
	// Adjust position for 'dot' position.
	const (
		dx = 3
		dy = 12
	)
	fx := fixed.I(x)
	fy := fixed.I(y)
	for i, t := range tokens(str) {
		clr := clr
		if i%2 == 1 {
			clr = red
		}
		text.Draw(target, t, bitmapfont.Gothic12r, fx.Round()+dx, fy.Round()+dy, clr)
		fx += font.MeasureString(bitmapfont.Gothic12r, t)
	}
}
