/*
> tinygo flash -target=pico ./L-chika/
> tinygo build -o L-chika.uf2 -target=pico -size short ./L-chika
   code    data     bss |   flash     ram
  15900     180    3184 |   16080    3364
ユーザボタン	　USER (GP24)
  LED  	青(GP25)
  LED	WS2812(GP23)
*/

package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	fmt.Printf("L-chika\n")
	led := machine.LED
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	for {
		fmt.Printf("LED off\n")
		led.Low()
		time.Sleep(time.Second / 10)
		fmt.Printf("LED on\n")
		led.High()
		time.Sleep(time.Second / 10)
	}
}
