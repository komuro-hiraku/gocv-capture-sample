package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	delay := flag.Int("delay", 0, "撮影までの待機時間（秒）")
	output := flag.String("output", "", "保存先の絶対パス（ファイル名含む）")
	flag.Parse()

	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("カメラを開けません:", err)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if *delay > 0 {
		fmt.Printf("⏳ %d秒待機中...\n", *delay)
		time.Sleep(time.Duration(*delay) * time.Second)
	}

	if ok := webcam.Read(&img); !ok || img.Empty() {
		fmt.Println("画像の取得に失敗しました")
		return
	}

	// 左右のみ切り出し（上下は元の高さを維持）
	width := img.Cols()
	height := img.Rows()
	centerX := width / 2
	cropWidth := width / 2
	
	// ROI（Region of Interest）を定義（左右のみカット）
	rect := image.Rect(centerX-cropWidth/2, 0, centerX+cropWidth/2, height)
	croppedImg := img.Region(rect)
	defer croppedImg.Close()
	
	// 切り出した画像をコピー
	finalImg := croppedImg.Clone()
	defer finalImg.Close()

	// 保存先の決定
	filename := "photo.jpg"
	if *output != "" {
		filename = *output
	} else {
		exePath, _ := os.Executable()
		filename = filepath.Join(filepath.Dir(exePath), filename)
	}

	gocv.IMWrite(filename, finalImg)
	fmt.Println("📸 撮影完了:", filename)
}
