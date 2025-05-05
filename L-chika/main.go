// > tinygo build -o L-chika.uf2 -target=pico -size short .
//    code    data     bss |   flash     ram
//   16640     180    3184 |   16820    3364
// > tinygo flash -target=pico -size short .

package main

import (
	"fmt"
	"machine"
	"time"
)

var LightOn time.Duration = 1
var LightOff time.Duration = 1

func main() {
	// ゲンジボタルの発光パターンを再現
	/* 点灯1秒 消灯1秒	長崎県五島列島
	LightOn = 1
	LightOff = 1
	*/
	/* 点灯2秒 消灯1秒	西日本と九州 */
	LightOn = 2
	LightOff = 1
	/* 点灯4秒 消灯1秒	東日本
	LightOn = 4
	LightOff = 1
	*/
	fmt.Printf("L-chika\n")
	led := machine.LED
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	for {
		fmt.Printf("LED off\n")
		led.Low()
		time.Sleep(time.Second * LightOff)
		fmt.Printf("LED on\n")
		led.High()
		time.Sleep(time.Second * LightOn)
	}
}
