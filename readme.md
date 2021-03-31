```shell
set GOOS=linux 
set GOARCH=amd64

go build -ldflags '-w -s'  -o ./cmd/go-iptv main.go

./cmd/upx ./cmd/go-iptv
set GOOS=windows 
set GOARCH=amd64
```