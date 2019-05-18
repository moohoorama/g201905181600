package draw

import (
	"bytes"
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/text/language"

	"github.com/hajimehoshi/go-inovation/ino/internal/assets"
	"github.com/hajimehoshi/go-inovation/ino/internal/fieldtype"
	"github.com/hajimehoshi/go-inovation/ino/internal/font"
	"github.com/hajimehoshi/go-inovation/ino/internal/input"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

var (
	imageItemFrame         *ebiten.Image
	imageItemMessageFrames = map[fieldtype.FieldType]*ebiten.Image{}
)

func init() {
	imageItemFrame, _ = ebiten.NewImage(32, 32, ebiten.FilterDefault)
	imageItemFrame.Fill(color.Black)
	ebitenutil.DrawRect(imageItemFrame, 2, 2, 28, 28, color.White)
}

func init() {
	title := map[fieldtype.FieldType]color.RGBA{
		fieldtype.FIELD_NONE:         {0xff, 0xff, 0xff, 0xff},
		fieldtype.FIELD_ITEM_POWERUP: {0x00, 0x76, 0x8a, 0xff},
		fieldtype.FIELD_ITEM_FUJI:    {0xe4, 0x32, 0x60, 0xff},
		fieldtype.FIELD_ITEM_TAKA:    {0x99, 0x5b, 0x00, 0xff},
		fieldtype.FIELD_ITEM_NASU:    {0x8a, 0x29, 0xd2, 0xff},
		fieldtype.FIELD_ITEM_OMEGA:   {0x25, 0xba, 0x18, 0xff},
		fieldtype.FIELD_ITEM_LIFE:    {0xe4, 0x32, 0x60, 0xff},
	}
	body := map[fieldtype.FieldType]color.RGBA{
		fieldtype.FIELD_ITEM_FUJI: {0xff, 0xc7, 0xd5, 0xff},
		fieldtype.FIELD_ITEM_TAKA: {0xed, 0xc0, 0x71, 0xff},
		fieldtype.FIELD_ITEM_NASU: {0xd3, 0xa0, 0xf9, 0xff},
	}

	for i, titleColor := range title {
		bodyColor, ok := body[i]
		if !ok {
			bodyColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
		}
		img, _ := ebiten.NewImage(256, 96, ebiten.FilterDefault)
		img.Fill(color.Black)
		ebitenutil.DrawRect(img, 1, 1, 254, 15, titleColor)
		ebitenutil.DrawRect(img, 1, 17, 254, 78, bodyColor)
		imageItemMessageFrames[i] = img
	}
}

func DrawItemFrame(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(imageItemFrame, op)
}

func DrawItemMessage(screen *ebiten.Image, item fieldtype.FieldType, y int, lang language.Tag) {
	frame, ok := imageItemMessageFrames[item]
	if !ok {
		frame = imageItemMessageFrames[fieldtype.FIELD_NONE]
	}
	w, _ := frame.Size()
	x := (ScreenWidth - w) / 2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(frame, op)

	str := item.ItemMessage(lang)
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		dx := (ScreenWidth - font.Width(line)) / 2
		dy := i * font.LineHeight
		if i == 0 {
			dy += 1
		} else {
			dy += (80 - font.LineHeight*(len(lines)-1)) / 2
		}
		clr := color.Black
		if i == 0 && (item == fieldtype.FIELD_ITEM_POWERUP ||
			item == fieldtype.FIELD_ITEM_FUJI ||
			item == fieldtype.FIELD_ITEM_TAKA ||
			item == fieldtype.FIELD_ITEM_NASU ||
			item == fieldtype.FIELD_ITEM_OMEGA ||
			item == fieldtype.FIELD_ITEM_LIFE) {
			clr = color.White
		}
		font.DrawText(screen, line, dx, y+dy, clr)
	}
}

var (
	images = map[string]*ebiten.Image{}
)

func LoadImages() error {
	for _, f := range []string{"ino", "msg_ja", "msg_en", "bg", "touch"} {
		b, err := assets.Asset("resources/images/color/" + f + ".png")
		if err != nil {
			return err
		}
		origImg, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			return err
		}
		images[f], _ = ebiten.NewImageFromImage(origImg, ebiten.FilterDefault)
	}
	return nil
}

func Draw(screen *ebiten.Image, key string, px, py, sx, sy, sw, sh int) {
	op := &ebiten.DrawImageOptions{}
	r := image.Rect(sx, sy, sx+sw, sy+sh)
	op.SourceRect = &r
	op.GeoM.Translate(float64(px), float64(py))
	screen.DrawImage(images[key], op)
}

func DrawTouchButtons(screen *ebiten.Image) {
	img := images["touch"]
	w, h := img.Size()
	w /= 4
	dx := 0
	dy := ScreenHeight - h
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.4)
	for _, i := range []int{0, 1, 3} {
		r := image.Rect(i*w, 0, (i+1)*w, h)
		op.SourceRect = &r
		op.GeoM.Reset()
		op.GeoM.Translate(float64(dx+i*w), float64(dy))
		screen.DrawImage(img, op)
	}
	// Render 'down' button
	op = &ebiten.DrawImageOptions{}
	r := image.Rect(2*w, 0, 3*w, h)
	op.SourceRect = &r
	op.GeoM.Translate(float64(dx+2*w), float64(dy))
	alpha := 0.0
	if input.Current().IsActionKeyPressed() {
		alpha = 0.4
	} else {
		alpha = 0.1
	}
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(img, op)
}
