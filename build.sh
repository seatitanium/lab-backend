export TZ=Asia/Shanghai
current=$(date '+%Y-%m-%dT%H:%M:%SZ+08:00')
env GOOS=linux GOARCH=amd64 go build -ldflags="-X seatimc/backend/common.LastBuiltAt=$current" -o lab