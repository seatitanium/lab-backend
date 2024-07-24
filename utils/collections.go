package utils

func HasKey[T interface{}](m map[string]T, key string) bool {
	_, ok := m[key]
	return ok
}

func LimitSlice[T interface{}](limit int, slice []T) []T {
	if limit > len(slice) {
		return slice
	}

	if limit <= 0 {
		return []T{}
	}

	return slice[:limit]
}
