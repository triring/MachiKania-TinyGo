// > tinygo build -o SixKeyPad.uf2 -target=pico -size short .
//    code    data     bss |   flash     ram
//    8232     108    3176 |    8340    3284
// > tinygo flash -target=pico -size short .

// 押されているキーを表示する。
// プログラムを書き込んだら、すぐに以下のコマンドを実行する。
// > tinygo monitor
// 0, 1, 0, 0, 1, 1,
// 0, 1, 0, 0, 1, 1,
// 0, 0, 0, 0, 0, 0,
// 0, 0, 0, 0, 0, 0,
// 0, 0, 0, 0, 0, 0,
// 0, 0, 0, 0, 0, 0,
// 1, 1, 1, 1, 0, 0,
// 1, 1, 1, 1, 0, 0,
// 1, 1, 1, 1, 0, 0,
// 0, 0, 0, 0, 0, 0,
// 1, 1, 0, 1, 0, 0,

package main

import (
	"machine"
	"time"
)

const (
	KeyMax = 6
)

// var gpioPins []machine.Pin
// rp2040
//gpioPins = []machine.Pin{
//	machine.GPIO4,  // up
//	machine.GPIO5,  // left
//	machine.GPIO6,  // down
//	machine.GPIO7,  // right
//	machine.GPIO27, // A
//	machine.GPIO28, // B
//}
// xiao-samd21 or xiao-rp2040
// gpioPins = []machine.Pin{
// 	machine.D10, // up
// 	machine.D9,  // left
// 	machine.D8,  // down
// 	machine.D7,  // right
// 	machine.D1,  // A
// 	machine.D2,  // B
// }

// Machikania type P Pico or YD-RP2040
var gpioPins = []machine.Pin{
	machine.GPIO8,  // up
	machine.GPIO9,  // left
	machine.GPIO21, // down
	machine.GPIO20, // right
	machine.GPIO22, // A(start)
	machine.GPIO26, // B(fire)
}

func main() {
	for _, p := range gpioPins {
		p.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	for {
		for _, p := range gpioPins {
			if p.Get() {
				print("0, ")
			} else {
				print("1, ")
			}
		}
		time.Sleep(time.Millisecond * 100)
		print("\n")
	}
}
