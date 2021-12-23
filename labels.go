package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"path/filepath"

	"github.com/fogleman/gg"
)

func drawLabel(baseImg image.Image, title string, fontSize float64) ([]byte, error) {
	dc := gg.NewContext(baseImg.Bounds().Size().X, baseImg.Bounds().Size().Y)
	dc.DrawImage(baseImg, 0, 0)

	addOverlay(dc, 20)

	textShadowColor := color.Black
	textColor := color.White
	fontPath := filepath.Join("fonts", "Open_Sans", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		return nil, fmt.Errorf("load font: %w", err)
	}

	textRightMargin := 60.0
	textTopMargin := 90.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin

	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignCenter)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignCenter)

	buf := &bytes.Buffer{}
	if err := dc.EncodePNG(buf); err != nil {
		return nil, fmt.Errorf("cannot encode: %w", err)
	}

	return buf.Bytes(), nil
}

func addOverlay(dc *gg.Context, margin float64) {
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()
}
