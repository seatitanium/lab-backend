package utils

func MustPanic(mustPanic error) {
	if mustPanic != nil {
		panic(mustPanic)
	}
}
