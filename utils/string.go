package utils

func AnyMatch(match string, strs ...string) bool {
	for _, str := range strs {
		if str == match {
			return true
		}
	}

	return false
}
