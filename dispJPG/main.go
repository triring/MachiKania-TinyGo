// > tinygo build -o dispJPG.uf2 --target=pico --size short .
//    code    data     bss |   flash     ram
//  116136    5008    3376 |  121144    8384
// > tinygo flash --target=pico --size short .

package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"machine"
	"strings"
	"time"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var (
	white  = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	red    = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	yellow = color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}
	green  = color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	blue   = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}
	black  = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
	indigo = color.RGBA{R: 0x3F, G: 0x48, B: 0xCC, A: 0xFF}
)

func main() {
	display := InitDisplay()

	// disp jpeg
	{
		img, err := jpeg.Decode(strings.NewReader(JimbeeJet320x240_jpg))
		if err != nil {
			log.Fatal(err)
		}

		w := img.Bounds().Dx()
		h := img.Bounds().Dy()
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				display.SetPixel(int16(x), int16(y), color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: 0xFF})
			}
		}
	}

	// text
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 28, 180, "JAPAN TRANSOCEAN AIR", indigo)
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 28, 196, "Jimbee Jet JA05RK", indigo)
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 28, 212, "Model Boeing 737-8Q3", indigo)
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 28, 228, "Camera Nikon Coolpix P1000", indigo)

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
	// backlight := machine.LCD_BACKLIGHT
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
