# GoCV Capture

Webカメラを使用した写真撮影ツール（Go + OpenCV）

## 機能

- Webカメラからの写真撮影
- 遅延撮影機能
- 左右切り出し機能（画角調整）
- Git pre-pushフックとの連携

## ビルド

```bash
go build -o bin/capture capture.go
```

## 実行

```bash
# 基本撮影
./bin/capture

# 2秒遅延撮影
./bin/capture -delay 2

# 保存先指定
./bin/capture -output /path/to/photo.jpg
```

## オプション

- `-delay`: 撮影までの待機時間（秒）
- `-output`: 保存先の絶対パス（ファイル名含む）

## Git Hook

pre-pushフック使用時、pushの前にcommit hashをファイル名とした写真が自動撮影されます。

```bash
# フックファイルを.git/hooksにコピー
cp hook/pre-push .git/hooks/pre-push
chmod +x .git/hooks/pre-push
```

## 依存関係

- Go 1.24.2+
- GoCV (gocv.io/x/gocv v0.41.0)
- OpenCV

## 画像処理

- 元画像の左右を中央から半分に切り出し
- 上下の高さは元のまま維持
- 結果として縦長のアスペクト比になる