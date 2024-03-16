package utils

import (
	"log"
	"slices"
)

func MustPanic(mustPanic error) {
	if mustPanic != nil {
		log.Fatal(mustPanic.Error())
	}
}

func NeedAuthorize(handlerName string) bool {
	return slices.Contains([]string{}, handlerName)
}
