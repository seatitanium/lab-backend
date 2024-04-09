package utils

import (
	"log"
	"time"
)

func MustPanic(mustPanic error) {
	if mustPanic != nil {
		log.Fatal(mustPanic.Error())
	}
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}
