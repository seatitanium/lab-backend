export TZ=Asia/Shanghai
env GOOS=linux GOARCH=amd64 go build -ldflags="-X seatimc/backend/common.LastBuiltAt=$(date +%Y-%m-%d_%H:%M:%S)+08:00" -o lab