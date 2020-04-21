package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

//Pic Return a matrix of uint8's to be drawn
func Pic(dx, dy int) [][]uint8 {

	fmt.Println("Image size")
	fmt.Printf("x:%d, y:%d\n", dx, dy)

	// allocate slice
	picture := make([][]uint8, dx)

	for i := 0; i < dx; i++ {
		// allocate slice
		row := make([]uint8, dy, dy)

		for j := 0; j < len(row); j++ {
			// Do some crazy things
			row[j] = uint8((i + j) / 2)
			row[len(row)-j-1] = row[j] ^ uint8(i*j)

		}
		picture[i] = row
	}
	return picture
}

// Show s
func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

// ShowImage s
func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

func main() {
	Show(Pic)
}
