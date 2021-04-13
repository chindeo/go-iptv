```shell
set GOOS=linux 
set GOARCH=amd64

CC=/home/go/src/github.com/snowlyg/go-iptv/go_ffmpeg/openwrt/bin/x86_64-openwrt-linux-gcc  go build -ldflags "-s -w" -o go_iptv go_ffmpeg/main.go

./cmd/upx ./cmd/go-iptv
set GOOS=windows 
set GOARCH=amd64
```