package utils

import (
	"log"
	"strings"
)

func MustPanic(mustPanic error) {
	if mustPanic != nil {
		log.Fatal(mustPanic.Error())
	}
}

func NeedAuthorize(handlerName string) bool {
	needs := Conf().NeedAuthorizeHandlers
	for i := 0; i < len(needs); i++ {
		if strings.HasPrefix(handlerName, needs[i]) {
			return true
		}
	}
	return false
}
