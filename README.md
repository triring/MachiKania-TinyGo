# MachiKania-TinyGo

Raspberry Pi Picoを利用してBASICプログラムを動作させることができるオープンプラットフォーム[MachiKania](https://github.com/machikania)の中の液晶搭載モデル [type P](http://www.ze.em-net.ne.jp/~kenken/machikania/typep.html) を[TinyGo](https://tinygo.org/)で動かしてみました。  

使用したMachiKania type P のスペックは以下の通りです。  

* 液晶	ILI9341搭載QVGA（320x240）SPI接続
* キー	上下左右、FIRE、START用計6個
* ブザー	圧電スピーカ 1

液晶は、SPI1に接続されており、標準的なili9341液晶ドライバで動きました。  
設定とピンの割当は以下の通りでした。  

| 項目 | 内容 |
|:-----|:-----|
| 液晶ドライバ | ili9341 |
| Raspiと液晶の接続 | SPI1 に接続 |
| 周波数設定 | 32000000 Hz |

| PIN    | 機能割当て |
|:-----|:-----|
| GPIO12 | SDI:Default Serial In Bus |
| GPIO15 | SDO:Default Serial Out Bus   1 for SPI communications MOSI |
| GPIO14 | SCK:Default Serial Clock Bus 1 for SPI communications SCLK |
| GPIO10 | LCD_DC |
| GPIO13 | LCD_CS |
| GPIO11 | LCD_RESET |


キーパッドの設定

| PIN    | キー   |
|:-------|:------ |
| GPIO4  | up     | 
| GPIO5  | left   | 
| GPIO6  | down   | 
| GPIO7  | right  | 
| GPIO27 | A      | 
| GPIO28 | B      | 

ブザーの設定

| PIN    | 機能         |
|:-------|:------------ |
| GPIO28 | PIEZO BUZZER | 
