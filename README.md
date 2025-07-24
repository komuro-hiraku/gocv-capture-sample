# Front cameraで自動的に写真とるやつ

## 動作環境

- go version go1.24.2 darwin/arm64
- macOS 14.7.6 (23H626)

## Install & Build

### Install opencv

```bash
brew install opencv
```

### Install go mod

```sh
go get gocv.io/x/gocv@latest
```

### Build

```sh
go build -o capture capture.go
```

### Initial Start

```sh
./bin/capture -delay 1 -output /Your/Path/hoge.jpg
```

:note: When you start the application for the first time, a dialog will appear requesting access to the camera. Please allow the application to access the camera.