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

// 经过 Unique 处理过的 slice 将保证数组中只有唯一元素满足 sameCondition
func Unique[T comparable](slice []T, sameCondition func(T, T) bool) []T {
	var newSlice []T

	for _, t := range slice {
		exists := false

		for _, nt := range newSlice {
			if sameCondition(nt, t) {
				exists = true
				break
			}
		}

		if !exists {
			newSlice = append(newSlice, t)
		}
	}

	return newSlice
}
