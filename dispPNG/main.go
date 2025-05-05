// > tinygo build -o dispPNG.uf2 --target=pico --size short .
// code    data     bss |   flash     ram
// 107696    2944    5496 |  110640    8440
// > tinygo flash --target=pico --size short .

package main

import (
	"image/color"
	"image/png"
	"log"
	"machine"
	"strings"
	"time"

	"tinygo.org/x/drivers/ili9341"
)

var (
	white  = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	red    = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	yellow = color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}
	green  = color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	blue   = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}
	black  = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
)

func main() {
	display := InitDisplay()

	// body
	// display.FillRectangle(20, 20, 280, 200, white)

	// lcd
	//	display.FillRectangle(25, 25, 270, 160, black)
	// gopher
	{
		img, err := png.Decode(strings.NewReader(Emergency_Exit_Gopher_200x200_png))
		// OK img, err := png.Decode(strings.NewReader(img80x80_png))
		// OK img, err := png.Decode(strings.NewReader(img128x128_png))
		// OK img, err := png.Decode(strings.NewReader(img144x144_png))
		// OK img, err := png.Decode(strings.NewReader(img160x160_png))
		// OK img, err := png.Decode(strings.NewReader(img180x180_png))
		// OK img, err := png.Decode(strings.NewReader(img200x200_png))
		// 200x200は、32byte幅で割り切れ、最後に＋が残った。変換ツールの終了条件をチェックすること。
		// 以下は、表示されなかっった。200x200が限界？
		// img, err := png.Decode(strings.NewReader(img220x220_png))

		// img, err := png.Decode(strings.NewReader(Emergency_Exit_Gopher_png))
		// img, err := png.Decode(strings.NewReader(Go_gopher_mascot_green_150x128_png))
		// img, err := png.Decode(strings.NewReader(Emergency_Exit_Gopher_200x200_png))
		// img, err := png.Decode(strings.NewReader(Emergency_Exit_Gopher_220x220_png))
		// img, err := png.Decode(strings.NewReader(Emergency_Exit_Gopher_240x240_png))

		if err != nil {
			log.Fatal(err)
		}

		w := img.Bounds().Dx()
		h := img.Bounds().Dy()
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				// display.SetPixel((320-int16(w))/2+int16(x), (240-int16(h))/2+int16(y), color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: 0xFF})
				display.SetPixel(60+int16(x), 20+int16(y), color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: 0xFF})
			}
		}
	}

	// speaker
	/*
		for i := int16(0); i < 4; i++ {
			display.FillRectangle(40+i*15, 190, 5, 20, black)
		}
	*/
	// buttons
	/*
		for i := int16(0); i < 3; i++ {
			display.FillRectangle(40+i*60, 15, 40, 5, blue)
		}
	*/
	// 5-way key
	//	tinydraw.FilledCircle(display, 260, 180, 20, blue)

	// text
	/*
		tinyfont.WriteLine(display, &freemono.Regular9pt7b, 30, 38, "Booting ", yellow)
		tinyfont.WriteLine(display, &freemono.Regular9pt7b, 30, 52, "MachiKania type P...", yellow)
	*/
	for {
		time.Sleep(time.Hour)
	}
}

func InitDisplay() *ili9341.Device {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.GPIO14, // Default Serial Clock Bus 1 for SPI communications SCLK
		SDO:       machine.GPIO15, // Default Serial Out Bus   1 for SPI communications MOSI
		SDI:       machine.GPIO12, // Default Serial In Bus    1 for SPI communications MISO
		Frequency: 32000000,
		Mode:      0,
	})

	// configure backlight
	//	backlight := machine.LCD_BACKLIGHT
	// backlight := machine.NoPin
	// backlight.Configure(machine.PinConfig{machine.PinOutput})
	display := ili9341.NewSPI(
		machine.SPI1,
		machine.GPIO10, // LCD_DC		dc := machine.GP10
		machine.GPIO13, // LCD_CS		cs := machine.GP13
		machine.GPIO11, // LCD_RESET	reset := machine.GP11
	)

	// configure display
	display.Configure(ili9341.Config{})

	//	backlight.High()

	display.SetRotation(ili9341.Rotation90)
	display.FillScreen(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF})

	return display
}
