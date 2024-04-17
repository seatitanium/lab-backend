package utils

func IsStrsHasEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return true
		}
	}
	return false
}
