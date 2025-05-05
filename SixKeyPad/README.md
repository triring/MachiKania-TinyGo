# SixKeyPad

MachiKania のキー入力をテストするサンプルです。  
MachiKania の6つのキーを監視し、押されているキーを表示します。動作テスト用として使ってください。  

## コンパイルと書き込み

1. 必要に応じて、以下のパッケージを追加インストールして下さい。すでにインストールされている場合は、不要です。  

```bash
go get tinygo.org/x/drivers
```

2. 以下のコマンドを実行してください。  

```bash
> tinygo build -o SixKeyPad.uf2 -target=pico -size short .
   code    data     bss |   flash     ram
  11944     108    3176 |   12052    3284
```

3. SixKeyPad.uf2 というファイルが生成されます。  
以下の手順で、このファイルをRaspberry Pi Picoに書き込んでください。

    1. Bootボタンを押しながらRaspberry Pi PicoにUSB給電すると[RPI-RP2]ドライブがマウントされ、書込みモードになります。  
    Pico互換ボードでは、Bootボタンを押した状態で、Resetボタンを押すと書込みモードになります。
    2. 先ほど作成した uf2ファイルを[RPI-RP2]ドライブにコピーします。  
    正常に書き込みができると強制アンマウントされ、プログラムがスタートします。

また、以下のコマンドを実行すると、コンパイル終了後に、直接、Raspberry Pi Picoに書込みをしてくれます。

```bash
> tinygo flash -target=pico -size short .
   code    data     bss |   flash     ram
   8232     108    3176 |    8340    3284
```

4. 動作確認は、プログラムの書き込みが完了してから、tinygo monitorを実行して下さい。  

* 1: 押されているキー
* 0: 押されていないキー

モニターの終了は、Ctrl-Cを押して下さい。  

```bash
> tinygo monitor
Connected to COM15. Press Ctrl-C to exit.

0, 1, 0, 0, 0, 0,
0, 1, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0,
0, 0, 1, 0, 0, 1,
0, 1, 1, 0, 0, 1,
0, 1, 1, 0, 0, 1,
0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0,
1, 0, 1, 0, 0, 1,
1, 0, 1, 0, 0, 1,
0, 0, 0, 1, 0, 0,
0, 0, 1, 0, 0, 0,
0, 0, 0, 0, 1, 0,
0, 0, 0, 0, 1, 0,
```

[Parent Directory](../README.md)