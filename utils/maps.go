package utils

func HasKey[T interface{}](m map[string]T, key string) bool {
	_, ok := m[key]
	return ok
}
