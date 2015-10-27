package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := i.Draw(x, y)
	return color.RGBA{v, v, 255, 255}
}

func (i Image) Draw(x, y int) uint8 {
	v := x * y * y
	return uint8(v)
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
