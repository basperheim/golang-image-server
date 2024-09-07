# Image Server Written in Golang

Simple image server written in Golang using the Gin framework.

## Run Gin Server

```bash
go run main.go
```

## Build Binaries

For ARM-based macOS Apple Silicon:

```bash
env GOOS=darwin GOARCH=arm64 go build -o cdn_server_macos_arm64
```

For x86 Linux:

```bash
env GOOS=linux GOARCH=amd64 go build -o cdn_server_linux_amd64
```

For 64-bit Windows:

```bash
env GOOS=windows GOARCH=amd64 go build -o cdn_server_windows_amd64.exe
```

## Usage

Make a cURL or Postman `GET` request like so to get image data:

```
curl -XGET "http://localhost:8282/cdn/test.jpeg"
```
