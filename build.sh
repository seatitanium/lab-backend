env GOOS=linux GOARCH=amd64 TZ=Asia/Shanghai
go build -ldflags="-X seatimc/backend/common.LastBuiltAt=$(date +%Y-%m-%d_%H:%M:%S)+08:00" -o lab