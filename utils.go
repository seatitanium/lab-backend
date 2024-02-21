package lab_backend

func MustPanic(mustPanic error) {
	if mustPanic != nil {
		panic(mustPanic)
	}
}
