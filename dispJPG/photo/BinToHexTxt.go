package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var filesize int64

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run BinToHexTxt.go <image_file>")
		os.Exit(1)
	}

	imagePath := os.Args[1]

	// ファイルの情報を確認する。
	file, err := os.Open(imagePath)
	defer file.Close()

	if fi, err := file.Stat(); err == nil {
		fmt.Printf("ファイル名: %s\n", fi.Name())
		fmt.Printf("ファイルサイズ(byte): %d\n", fi.Size())
		fmt.Printf("モード: %s\n", fi.Mode())
		fmt.Printf("ディレクトリ? :%t\n", fi.IsDir())
		// var size = fi.Size()
		// fmt.Printf(reflect.TypeOf(size))
		filesize = fi.Size()
	}
	fmt.Printf("ファイルサイズ(byte): %d\n", filesize)
	file.Close()

	// 画像ファイルとして、開き直す。
	data, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		os.Exit(1)
	}

	hexString := strings.ToUpper(hex.EncodeToString(data)) // 大文字に変換

	outputFileName := strings.Replace(filepath.Base(imagePath), ".", "_", -1) + ".txt" // バイナリをテキスト形式で出力
	//	outputFileName := filepath.Base(imagePath) + ".txt" // バイナリをテキスト形式で出力
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating Data file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()
	LastLine := false
	LineCount := 0
	for i := 0; i < len(hexString); i += 64 { // 32個の16進数（64文字）ごとに改行
		end := i + 64
		if end > len(hexString) {
			end = len(hexString)
			LastLine = true
		}
		line := hexString[i:end]

		var csvLine []string
		for j := 0; j < len(line); j += 2 {
			csvLine = append(csvLine, line[j:j+2])
		}
		var oneline string
		//	oneline += fmt.Sprintf("/* %d */", LineCount)
		oneline += "\t\"\\x"
		oneline += strings.Join(csvLine, "\\x")
		// 行末で文字列の連結を行うために'+'を書き加える。
		// しかし、最終行には不要なので、行末か否かの判別して、出力を切り替える。
		if LastLine != true {
			oneline += "\" +"
		} else {
			oneline += "\""
		}
		LineCount++

		_, err = fmt.Fprintln(outputFile, oneline)
		//		_, err = fmt.Fprintln(outputFile, strings.Join(csvLine, "\\"))
		if err != nil {
			fmt.Println("Error writing to Data file:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("行数: %d\n", LineCount)
	fmt.Println("Successfully converted to", outputFileName)
}
