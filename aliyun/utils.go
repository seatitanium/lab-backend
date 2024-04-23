package aliyun

func GetIoOptimized(boolean bool) string {
	if boolean {
		return "optimized"
	} else {
		return "none"
	}
}
