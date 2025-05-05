// > tinygo build -o buzzer.uf2 -target=pico -size short .
//    code    data     bss |   flash     ram
//   11944     108    3176 |   12052    3284
// > tinygo flash -target=pico -size short .

package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

type note struct {
	tone     float64
	duration float64
}

func main() {
	// bzrPin := machine.GPIO22 // PIEZO BUZZER Cytron Maker nano
	bzrPin := machine.GPIO28 // PIEZO BUZZER machikania typep

	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr := buzzer.New(bzrPin)

	notes := []note{
		{buzzer.C3, buzzer.Quarter},
		{buzzer.Rest, buzzer.Eighth},
		{buzzer.D3, buzzer.Eighth},
		{buzzer.E3, buzzer.Quarter},
		{buzzer.Rest, buzzer.Eighth},
		{buzzer.C3, buzzer.Eighth},
		{buzzer.E3, buzzer.Quarter},
		{buzzer.C3, buzzer.Quarter},
		{buzzer.E3, buzzer.Half},
	}

	for _, n := range notes {
		bzr.Tone(n.tone, n.duration)
		time.Sleep(10 * time.Millisecond)
	}

	for {
		time.Sleep(time.Hour)
	}
}
