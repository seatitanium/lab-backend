n=$(cat buildnumber)
nn=$((n+1))
env GOOS=linux GOARCH=amd64 TZ=Asia/Shanghai
go build -ldflags="-X seatimc/backend/common.LastBuiltAt=$nn.$(date +%Y-%m-%d_%H:%M:%S)+08:00" -o lab
echo $nn > buildnumber