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

func ParseTime(layout string, str string) (time.Time, error) {
	res, err := time.Parse(time.RFC3339, str)

	if err == nil {
		return res, nil
	}

	return time.Parse(layout, str)
}

func ParseTimeRFC3339(str string) (time.Time, error) {
	return ParseTime("2006-01-02T15:04:05Z", str)
}

func ParseTimeRFC3339Milli(str string) (time.Time, error) {
	return ParseTime("2006-01-02T15:04:05.999Z07:00", str)
}
