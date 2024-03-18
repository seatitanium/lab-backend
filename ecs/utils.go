package ecs

func GetIoOptimized(boolean bool) string {
	if boolean {
		return "optimized"
	} else {
		return "none"
	}
}
