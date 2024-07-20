package utils

func HasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
