# buzzer

MachiKania のブザーを鳴らします。   
動作テスト用として使ってください。  

## コンパイルと書き込み

1. 必要に応じて、以下のパッケージを追加インストールして下さい。すでにインストールされている場合は、不要です。  

```bash
go get tinygo.org/x/drivers/buzzer
```

2. 以下のコマンドを実行してください。  

```bash
> tinygo build -o buzzer.uf2 -target=pico -size short .
   code    data     bss |   flash     ram
  11944     108    3176 |   12052    3284
```

3. buzzer.uf2 というファイルが生成されます。  
以下の手順で、このファイルをRaspberry Pi Picoに書き込んでください。

    1. Bootボタンを押しながらRaspberry Pi PicoにUSB給電すると[RPI-RP2]ドライブがマウントされ、書込みモードになります。  
    Pico互換ボードでは、Bootボタンを押した状態で、Resetボタンを押すと書込みモードになります。
    2. 先ほど作成した uf2ファイルを[RPI-RP2]ドライブにコピーします。  
    正常に書き込みができると強制アンマウントされ、プログラムがスタートします。

また、以下のコマンドを実行すると、コンパイル終了後に、直接、Raspberry Pi Picoに書込みをしてくれます。

```bash
> tinygo flash -target=pico -size short .
   code    data     bss |   flash     ram
  11944     108    3176 |   12052    3284
```

[Parent Directory](../README.md)