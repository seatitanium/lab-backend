env GOOS=linux GOARCH=amd64
go build -o lab -ldflags='-X common.LastBuiltAt="date +%Y-%m-%d_%H:%M:%S"'