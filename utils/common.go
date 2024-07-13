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

func ParseTimeATBZ(str string) (time.Time, error) {
	res, err := time.Parse(time.RFC3339, str)

	if err == nil {
		return res, nil
	}

	return time.Parse("2006-01-02T15:04Z", str)
}
